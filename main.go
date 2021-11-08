package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/popsicleposse/dotdotbot/config"
	"github.com/popsicleposse/dotdotbot/contracts/dotdotdots"
	"github.com/popsicleposse/dotdotbot/model"
)

const ()

var (
	configPath = flag.String("config", "config.json", "defines the path to the config")

	conf *config.Config

	database *gorm.DB

	keyStore *keystore.KeyStore
)

func init() {
	file, err := os.Open(*configPath)
	if err != nil {
		// something wrong with the config path (path doesn't exist?), just fail
		log.Fatalln(err)
	}

	c, err := config.ReadConfig(file)

	if err != nil {
		// failed to read the config, not much we can do to recover
		log.Fatalln(err)
	}

	conf = c // set the config to the one read
	store := keystore.NewKeyStore(c.KeystoreConf.Path, keystore.StandardScryptN, keystore.StandardScryptP)

	if store == nil {
		log.Fatalln("store is nil")
	} else if len(store.Accounts()) == 0 {
		account, err := store.NewAccount(conf.KeystoreConf.Password)

		if err != nil {
			log.Fatalln("failed to create new account in keystore:", err)
		}

		log.Printf("created new account %s, fund this wallet with Ethereum.\n", account.Address.Hex())
	}

	keyStore = store

	db, err := gorm.Open(sqlite.Open(conf.DBPath), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// create the migration for the database
	db.AutoMigrate(&model.QueuedTxn{})
	db.AutoMigrate(&model.FinalTxn{})

	database = db
}

func ethToWei(eth float64) *big.Int {
	n, _ := new(big.Float).Mul(big.NewFloat(eth), big.NewFloat(params.GWei)).Int(nil)
	n = n.Mul(n, big.NewInt(params.GWei))
	return n
}

func main() {
	// attempt to create a new ethereum client using the host we have
	client, err := ethclient.Dial(conf.Host)

	if err != nil {
		// could create the new client, is the host down? fail with the reason
		log.Fatalln(err)
	}

	var (
		mintContractAddress = common.HexToAddress(conf.MintContract)
		successfulMints     = 0
		// retries         = 1
	)

	chainId, err := client.ChainID(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	mintContract, err := dotdotdots.NewDotdotdots(mintContractAddress, client)

	if err != nil {
		// Could not create the contract, oh well
		log.Fatalln(err)
	}

	name, err := mintContract.Name(nil)

	if err != nil {
		log.Fatalln(err)
	}

	contractAbi, err := dotdotdots.DotdotdotsMetaData.GetAbi()

	if err != nil {
		log.Fatalln(err)
	}

	/*

		Honestly what i need to do is have a block subscriber that'll scan for transactions on the block for searching for
		the setSaleIsActive or whatever...
	*/

	// literally just want to run it as long as we can
	for {

		selectedAccount := keyStore.Accounts()[0]

		// prepare the account to send a transaction
		keyStore.Unlock(selectedAccount, conf.KeystoreConf.Password)
		// check if the sale is
		activeSale, err := mintContract.SaleIsActive(&bind.CallOpts{})

		totalSupply, _ := mintContract.TotalSupply(&bind.CallOpts{})
		maxSupply, _ := mintContract.MAXSUPPLY(nil)
		if err != nil {
			// print the error
			log.Println(err)

			// max = 4460, totalSupply = 4360, we don't want to mint at 4460 because transaction will fail
			// we need room for the amount we want to mint
			// so max = 4454 = totalSupply = 4450 there would be enough room for 6 mints
			// but with 4455 +  6 that would put us above the max supply of 4460 (4461).
		} else if activeSale {
			if totalSupply.Int64() < maxSupply.Int64()-int64(conf.Mint.MintTarget) {
				var pendingTxns []model.QueuedTxn
				var newTxns []common.Hash

				database.Where(&model.QueuedTxn{Pending: true}).Find(&pendingTxns)

				if len(pendingTxns) > 0 {

					for _, txn := range pendingTxns {

						val, _ := new(big.Int).SetString(txn.Value, 10)
						data := hexutil.MustDecode(txn.Data)

						msg := ethereum.CallMsg{
							From:  selectedAccount.Address,
							To:    &mintContractAddress,
							Value: val,
							Data:  data,
						}
						gasEstimate, err := client.EstimateGas(context.Background(), msg)
						price, _ := client.SuggestGasPrice(context.Background())

						if err != nil {
							log.Println(err)
						} else {
							tx, err := mintContract.Mint(&bind.TransactOpts{
								Value: val,
								Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
									return keyStore.SignTx(selectedAccount, tx, chainId)
								},
								GasPrice: big.NewInt(1).Mul(price, big.NewInt(int64(conf.Mint.GasMultiplier))),
								GasLimit: gasEstimate * conf.Mint.GasMultiplier,
								Nonce:    big.NewInt(int64(txn.Nonce)),
							}, big.NewInt(int64(txn.Amount)))

							if err != nil {
								log.Println(err)
							} else {

								newTxns = append(newTxns, tx.Hash())
								// transaction no longer pending, update it.
								database.Where("hash = ?", txn.Hash).Delete(&txn)

								fmt.Println("successfully replayed the transaction, new txn hash:", tx.Hash().String())
							}

						}
					}
				} else {

					target := conf.Mint.MintTarget
					perTransaction := conf.Mint.MintPerTransaction
					transactionCount := int(math.Ceil(float64(target) / float64(perTransaction)))

					for i := 0; i < transactionCount; i++ {

						mintCount := int(math.Min(float64((perTransaction)), float64(target-uint64(transactionCount))))
						data, _ := contractAbi.Pack("mint", big.NewInt(int64(mintCount)))

						msg := ethereum.CallMsg{
							From:  selectedAccount.Address,
							To:    &mintContractAddress,
							Value: ethToWei(conf.Mint.Price * float64(mintCount)),
							Data:  data,
						}
						gasEstimate, err := client.EstimateGas(context.Background(), msg)

						price, _ := client.SuggestGasPrice(context.Background())
						if err != nil {
							log.Println(err)
						} else {
							transaction, err := mintContract.Mint(&bind.TransactOpts{
								From:      selectedAccount.Address,
								Value:     msg.Value,
								GasPrice:  big.NewInt(1).Mul(price, big.NewInt(int64(conf.Mint.GasMultiplier))),
								GasLimit:  gasEstimate,    // set to a minimum
								GasTipCap: ethToWei(0.01), // tip the miner when submitting the transaction
								Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
									return keyStore.SignTx(selectedAccount, tx, chainId)
								},
							}, big.NewInt(int64(mintCount)))

							if err != nil {
								log.Println(err)
							} else {
								newTxns = append(newTxns, transaction.Hash())
							}
						}

					}

				}

				for _, txnHash := range newTxns {
					receipt, err := client.TransactionReceipt(context.Background(), txnHash)

					for err != nil {
						time.Sleep(500 * time.Millisecond)

						log.Println("no receipt yet. retrying.")
						receipt, err = client.TransactionReceipt(context.Background(), txnHash)
					}

					txn, pending, err := client.TransactionByHash(context.Background(), txnHash)

					for err != nil || pending {
						time.Sleep(500 * time.Millisecond)
						log.Println("no log yet. retrying.")
						txn, pending, err = client.TransactionByHash(context.Background(), txnHash)
					}

					database.Create(
						&model.FinalTxn{
							ProjectName: name,
							Status:      uint(receipt.Status),
							Hash:        txnHash.String()[2:],
							Value:       txn.Value().String(),
							Gas:         txn.Gas(),
							GasPrice:    txn.GasPrice().String(),
							Cost:        txn.Cost().String(),
						})
				}

				break
			} else {
				time.Sleep(850 * time.Millisecond)
				log.Println("waiting for supply increase...")
			}

		} else {
			// if the sale is not active we want to place a transaction that will remain in pending util we update it.

			target := conf.Mint.MintTarget
			perTransaction := conf.Mint.MintPerTransaction

			transactionCount := int(math.Ceil(float64(target) / float64(perTransaction)))

			var queuedTxns []model.QueuedTxn

			database.Where(&model.QueuedTxn{}).Find(&queuedTxns)

			// dropped := false
			// for _, queuedTxn := range queuedTxns {
			// 	txn, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(queuedTxn.Hash))

			// 	if err != nil {
			// 	}
			// }

			if len(queuedTxns) < int(transactionCount) {
				log.Println("queueing", transactionCount, "transactions")

				totalQueued := 0
				for i := 0; i < transactionCount; i++ {

					mintCount := int(math.Min(float64((perTransaction)), float64(target-uint64(totalQueued))))
					// grab the latest nonce we should be using...this is the number of pending transactions or simply the next
					// nonce (final transaction count + 1)
					nonce, err := client.PendingNonceAt(context.Background(), selectedAccount.Address)
					if err != nil {
						log.Println("error:", err)
					} else {
						// calculate the value
						value := ethToWei(conf.Mint.Price * float64(mintCount))

						transaction, err := mintContract.Mint(&bind.TransactOpts{
							Value:    value,
							GasPrice: big.NewInt(8), // set to a low gas price
							GasLimit: 530000,        // set to a minimum
							Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
								return keyStore.SignTx(selectedAccount, tx, chainId)
							},
							// NoSend: true,
							Nonce: big.NewInt(int64(nonce)),
						}, big.NewInt(int64(mintCount)))

						if err != nil {
							log.Println(err)
						} else {
							dbtx := &model.QueuedTxn{
								ProjectName: name,
								Hash:        transaction.Hash().String()[2:],
								Pending:     true,
								Status:      0,
								Data:        hexutil.Encode(transaction.Data()),
								Value:       transaction.Value().String(),
								Amount:      uint64(mintCount),
								Nonce:       transaction.Nonce(),
							}

							totalQueued += mintCount

							database.Create(dbtx)
						}

					}
				}

				log.Println("successfully queued", transactionCount, "transactions...awaiting mint.")
			} else {
				log.Println(len(queuedTxns), "in queue...awaiting sale.")
				time.Sleep(850 * time.Millisecond) // sleep for 1 second to reduce disk reads...
			}

		}

		keyStore.Lock(keyStore.Accounts()[0].Address) // relock the account
	}

	log.Printf("successfully minted %d NFT(s). target reached. shutting down\n", successfulMints)
}
