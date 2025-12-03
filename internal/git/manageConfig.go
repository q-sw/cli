package git

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func SwitchConfig(configName string) {
	configs := utils.FetchFiles("gitConfigPath")
	configPath := viper.GetString("gitConfigPath")
	homeDir := utils.GetHomeDir()

	var choice string
	if configName == "" {
		choice = filepath.Join(configPath, utils.List(configs))

	} else {
		choice = filepath.Join(configPath, configName)
	}

	err := os.Remove(filepath.Join(homeDir, ".gitconfig"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = os.Symlink(choice, filepath.Join(homeDir, ".gitconfig"))
	if err != nil {
		log.Println("error to create symlink")
		os.Exit(1)
	}
	log.Printf("%s is enabled\n", choice)
}

func GetCurrentConfig() {
	config, err := os.Readlink(filepath.Join(utils.GetHomeDir(), ".gitconfig"))
	if err != nil {
		log.Println("error to read link info")
	}
	_, c := filepath.Split(config)

	fmt.Println(c)
}
