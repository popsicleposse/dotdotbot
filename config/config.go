package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// lol, small config
type Config struct {
	Host         string   `json:"host"`
	KeystoreConf KeyStore `json:"keystore"`
	MintContract string   `json:"mintContract"`
	BotContract  string   `json:"botContract"`
	MintCount    uint     `json:"mintCount"`
	MaxMintCount uint     `json:"maxMintCount"`
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
