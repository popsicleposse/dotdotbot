package main

import (
	"flag"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/popsicleposse/dotdotbot/config"
	"github.com/popsicleposse/dotdotbot/contracts/dotdotbot"
	"github.com/popsicleposse/dotdotbot/contracts/kaiju"
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
		botContractAddress  = common.HexToAddress(conf.BotContract)
	)

	mintContract, err := kaiju.NewKaiju(mintContractAddress, client)

	if err != nil {
		// Could not create the contract, oh well
		log.Fatalln(err)
	}

	botContract, err := dotdotbot.NewDotdotbot(botContractAddress, client)

	if err != nil {
		// Could not create the contract, oh well
		log.Fatalln(err)
	}
	// literally just want to run it as long as we can
	for {
		// check if the sale is
		activeSale, err := mintContract.SaleActive(&bind.CallOpts{})

		if err != nil {
			// print the error
			log.Println(err)
		} else if activeSale {

			keyStore.Unlock(keyStore.Accounts()[0], conf.KeystoreConf.Password)
			// try to mint. The node will need to have a private key (wallet) associated with it! this can be set up through an encrypted keystore
			// or an account attached through `geth account`
			// See: https://geth.ethereum.org/docs/interface/managing-your-accounts
			transaction, err := botContract.TryMint(&bind.TransactOpts{
				From: keyStore.Accounts()[0].Address,
				Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
					return keyStore.SignTx(keyStore.Accounts()[0], tx, tx.ChainId())
				},
			}, big.NewInt(int64(conf.MintCount)))

			if err != nil {
				log.Println("failed transaction:", err)
			} else {
				log.Println(transaction.Hash().Hex())
			}

			keyStore.Lock(keyStore.Accounts()[0].Address)

		}
	}
}
