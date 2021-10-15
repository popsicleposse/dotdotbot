package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/popsicleposse/dotdotbot/config"
	"github.com/popsicleposse/dotdotbot/contracts/dotdotdots"
)

const ()

var (
	configPath = flag.String("config", "config.json", "defines the path to the config")

	conf *config.Config

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
		retries         = 1
	)

	mintContract, err := dotdotdots.NewDotdotdots(mintContractAddress, client)

	if err != nil {
		// Could not create the contract, oh well
		log.Fatalln(err)
	}

	// botContract, err := dotdotbot.NewDotdotbot(botContractAddress, client)

	// if err != nil {
	// 	// Could not create the contract, oh well
	// 	log.Fatalln(err)
	// }

	// abi, err := dotdotbot.DotdotbotMetaData.GetAbi()

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// literally just want to run it as long as we can
	for uint64(successfulMints) < conf.Mint.MaxMintCount {
		// check if the sale is
		activeSale, err := mintContract.SaleIsActive(&bind.CallOpts{})

		if err != nil {
			// print the error
			log.Println(err)
		} else if activeSale {

			selectedAccount := keyStore.Accounts()[0]
			keyStore.Unlock(selectedAccount, conf.KeystoreConf.Password)

			// try to mint. The node will need to have a private key (wallet) associated with it! this can be set up through an encrypted keystore
			// or an account attached through `geth account`
			// See: https://geth.ethereum.org/docs/interface/managing-your-accounts

			var (
				transaction *types.Transaction
				err         error
			)

			if !conf.MintWithContract {
				transaction, err = mintContract.Mint(&bind.TransactOpts{
					From: selectedAccount.Address,
					Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
						return keyStore.SignTx(selectedAccount, tx, tx.ChainId())
					},
					Value:  ethToWei(conf.Mint.Price * float64(conf.Mint.MintCount)),
					NoSend: true,
				}, big.NewInt(int64(conf.Mint.MintCount)))

				if err != nil {
					log.Println(err)
				} else {
					gas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
						From:  selectedAccount.Address,
						To:    transaction.To(),
						Data:  transaction.Data(),
						Value: transaction.Value(),
					})

					if err != nil {
						log.Println(err)
					} else {
						transaction, err = mintContract.Mint(&bind.TransactOpts{
							From: selectedAccount.Address,
							Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
								return keyStore.SignTx(selectedAccount, tx, big.NewInt(5))
							},
							Value:    ethToWei(conf.Mint.Price * float64(conf.Mint.MintCount)),
							GasLimit: gas * conf.Mint.GasMultiplier,
							GasPrice: big.NewInt(8),
						}, big.NewInt(int64(conf.Mint.MintCount)))

						if err != nil {
							log.Println(err)
						} else {

							log.Printf("submitted transaction %s, gas used: %d, cost: %s. awaiting receipt\n", transaction.Hash().String(), transaction.Gas(), transaction.Cost().String())

							time.Sleep(5 * time.Second)
							log.Printf("resending the transaction with new gas...")

							_, pending, _ := client.TransactionByHash(context.Background(), transaction.Hash())

							if pending {
								price, _ := client.SuggestGasPrice(context.Background())

								transaction.GasPrice().Add(price, big.NewInt(0))

								log.Println(client.SendTransaction(context.Background(), transaction))
							}

							receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())

							for err != nil {
								receipt, err = client.TransactionReceipt(context.Background(), transaction.Hash())
								time.Sleep(50 * time.Millisecond)
							}

							if receipt.Status >= 1 {
								successfulMints++
								log.Printf("submitted transaction %s was successful.\n", transaction.Hash())
							} else {
								retries--

							}
						}

					}
				}

			}

			keyStore.Lock(keyStore.Accounts()[0].Address)
		} else {
			log.Println("sale not active.")
			time.Sleep(50 * time.Millisecond)
		}

		if retries == 0 {
			log.Println("retries reached zero. shutting down.")
			break
		}
	}

	log.Printf("successfully minted %d NFT(s). target reached. shutting down\n", successfulMints)
}
