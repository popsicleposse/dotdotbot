package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// lol, small config
type Config struct {
	Host         string      `json:"host"`
	KeystoreConf KeyStore    `json:"keystore"`
	MintContract string      `json:"mintContract"`
	BotContract  string      `json:"botContract"`
	Mint         MintDetails `json:"mint"`
	DBPath       string      `json:"dbPath"`
	// Do we try to mint on chain with our own deployed contract?
	MintWithContract bool `json:"mintWithContract"`
}

type MintDetails struct {
	// We will multiply the gas to use by this much to attempt to ensure a mint.
	GasMultiplier uint64 `json:"gasMultiplier"`
	// The price of the mint in question, only used when not minting with contract - the price gets set on the contract
	Price float64 `json:"price"`
	// The number we should attempt to mint in total
	MintTarget uint64 `json:"mintTarget"`
	// the number of mints in a transaction
	MintPerTransaction uint64 `json:"mintPerTransaction"`
}

type KeyStore struct {
	Path     string
	Password string
}

func ReadConfig(r io.Reader) (*Config, error) {
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	var config Config // define an empty config struct

	if err = json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil // return a nil error and the pointer to our parsed config
}
