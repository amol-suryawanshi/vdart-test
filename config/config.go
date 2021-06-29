package config

import (
	"encoding/json"
	"io/ioutil"
)

//ServerConfig is in-memory configuration
var ServerConfig *configuration

type configuration struct {
	BaseURL    string     `json:"baseURL"`
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Currency string `json:"currency"`
	Symbol   string `json:"symbol"`
}

//ReadConfig reads configuration from given path
func ReadConfig(configFilePath string) error {
	ServerConfig = &configuration{}
	byteArr, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteArr, ServerConfig)
	if err != nil {
		return err
	}
	return nil
}

//LoadApplication loads config and logger
func LoadApplication(configFilePath string) error {
	err := ReadConfig(configFilePath)
	if err != nil {
		return err
	}
	return nil
}
