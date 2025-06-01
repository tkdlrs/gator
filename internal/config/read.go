package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	// get the path to the config
	pathToConfig, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	fmt.Println(pathToConfig)
	// Read the config JSON file
	data, err := os.ReadFile(pathToConfig)
	if err != nil {
		fmt.Println("error reading data.")
		return Config{}, err
	}
	fmt.Println("The raw data is: ", data)
	//
	configResp := Config{}
	err = json.Unmarshal(data, &configResp)
	if err != nil {
		return Config{}, err
	}
	fmt.Println("The configResp is: ", configResp)
	fmt.Println(configResp.DbUrl, "Could you be... My db url?")
	fmt.Println(configResp.CurrentUserName, "Current User Name?")
	//
	return configResp, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = "/.gatorconfig.json"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + configFileName, nil
}

func (cfg Config) SetUser() {
	write(cfg)
}

func write(cfg Config) error {
	//
	pathToConfig, err := getConfigFilePath()
	if err != nil {
		return err
	}
	//
	cfg.CurrentUserName = "levi"
	// prep data for write
	prepUpdatedCfg, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	// write data
	err = os.WriteFile(pathToConfig, prepUpdatedCfg, 0666)
	if err != nil {
		return err
	}
	//
	return nil
}
