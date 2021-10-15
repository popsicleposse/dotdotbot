package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

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
	db.AutoMigrate(&model.StoredTxn{})

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
		// botContractAddress  = common.HexToAddress(conf.BotContract)

		successfulMints = 0
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

	// contractAbi, err := dotdotdots.DotdotdotsMetaData.GetAbi()

	if err != nil {
		log.Fatalln(err)
	}

	// literally just want to run it as long as we can
	for uint64(successfulMints) < conf.Mint.MaxMintCount {

		selectedAccount := keyStore.Accounts()[0]

		// prepare the account to send a transaction
		keyStore.Unlock(selectedAccount, conf.KeystoreConf.Password)
		// check if the sale is
		activeSale, err := mintContract.SaleIsActive(&bind.CallOpts{})

		if err != nil {
			// print the error
			log.Println(err)
		} else if activeSale {
			var pendingTxns []model.StoredTxn
			database.Where(&model.StoredTxn{Pending: true}).Find(&pendingTxns)

			for _, txn := range pendingTxns {
				tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(txn.Hash))

				if err != nil {
					log.Println(err)
				} else if pending {

					msg := ethereum.CallMsg{
						From:  selectedAccount.Address,
						To:    &mintContractAddress,
						Value: tx.Value(),
						Data:  tx.Data(),
					}
					gasEstimate, err := client.EstimateGas(context.Background(), msg)

					if err != nil {
						log.Println(err)
					} else {
						tx, err := mintContract.Mint(&bind.TransactOpts{
							Value: tx.Value(),
							Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
								return keyStore.SignTx(selectedAccount, tx, chainId)
							},
							GasLimit: gasEstimate * conf.Mint.GasMultiplier,
							Nonce:    big.NewInt(int64(txn.Nonce)),
						}, big.NewInt(int64(conf.Mint.MintCount)))

						if err != nil {
							log.Println(err)
						} else {

							database.Table("stored_txns").Where("hash = ?", txn.Hash).Update("pending", false)

							fmt.Println("successfully replayed the transaction", tx)
						}

					}

				} else {
					// transaction no longer pending, update it.
					database.Table("stored_txns").Where("hash = ?", txn.Hash).Update("pending", false)
				}
			}
		} else {
			// if the sale is not active we want to place a transaction that will remain in pending util we update it.

			// grab the latest nonce we should be using...this is the number of pending transactions or simply the next
			// nonce (final transaction count + 1)
			nonce, err := client.PendingNonceAt(context.Background(), selectedAccount.Address)
			if err != nil {
				log.Println(err)
			} else {

				// calculate the value
				value := ethToWei(conf.Mint.Price * float64(conf.Mint.MintCount))

				transaction, err := mintContract.Mint(&bind.TransactOpts{
					Value:    value,
					GasPrice: big.NewInt(8), // set to a low gas price
					GasLimit: 530000,        // set to a minimum
					Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
						return keyStore.SignTx(selectedAccount, tx, chainId)
					},
					// NoSend: true,
					Nonce: big.NewInt(int64(nonce)),
				}, big.NewInt(int64(conf.Mint.MintCount)))

				if err != nil {
					log.Println(err)
				} else {

					fmt.Println(transaction.Hash())
					dbtx := &model.StoredTxn{
						Hash:    transaction.Hash().String(),
						Pending: true,
						Status:  0,
						Data:    hexutil.Encode(transaction.Data()),
						// Value:   transaction.Value(),
						Nonce: transaction.Nonce(),
					}

					database.Create(dbtx)
				}

			}

			// transaction, err := mintContract.Mint(&bind.TransactOpts{
			// 	From: selectedAccount.Address,
			// 	Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// 		return keyStore.SignTx(selectedAccount, tx, big.NewInt(5))
			// 	},
			// 	Value: ethToWei(conf.Mint.Price * float64(conf.Mint.MintCount)),

			// 	GasPrice: big.NewInt(8),
			// }, big.NewInt(int64(conf.Mint.MintCount)))

		}

		keyStore.Lock(keyStore.Accounts()[0].Address) // relock the account
	}

	log.Printf("successfully minted %d NFT(s). target reached. shutting down\n", successfulMints)
}
