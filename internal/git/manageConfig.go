package git

import (
	"log"
	"os"
	"path/filepath"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func SwitchConfig(configName string) error {
	configs, err := utils.FetchFiles("gitConfigPath")
	if err != nil {
		return err
	}
	configPath := viper.GetString("gitConfigPath")

	homeDir, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	var choice string
	if configName == "" {
		choice = filepath.Join(configPath, utils.List(configs))

	} else {
		choice = filepath.Join(configPath, configName)
	}

	err = os.Remove(filepath.Join(homeDir, ".gitconfig"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = os.Symlink(choice, filepath.Join(homeDir, ".gitconfig"))
	if err != nil {
		log.Println("error to create symlink")
		return err
	}
	log.Printf("%s is enabled\n", choice)
	return nil
}

func GetCurrentConfig() (string, error) {
	homeDir, err := utils.GetHomeDir()
	if err != nil {
		return "", err
	}
	config, err := os.Readlink(filepath.Join(homeDir, ".gitconfig"))
	if err != nil {
		return "", err
	}
	_, c := filepath.Split(config)

	return c, nil
}
