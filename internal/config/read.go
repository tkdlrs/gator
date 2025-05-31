package config

import (
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
	fmt.Println("The data is: ", data)
	//

	//
	return Config{}, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + configFileName, nil
}

func (c Config) SetUser() {

}

func write(cfg Config) error {
	return fmt.Errorf("ERROR: wrong write.")
}
