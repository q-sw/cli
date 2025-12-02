package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/viper"
)

func FetchFiles(configPath string) []list.Item {
	c := viper.GetString(configPath)

	if c == "" {
		fmt.Printf("Error: %s is not set in your config file.\n", configPath)
		os.Exit(1)
	}

	configFiles, err := os.ReadDir(c)
	if err != nil {
		fmt.Printf("Error reading directory '%s': %v\n", c, err)
		os.Exit(1)
	}

	var items []list.Item
	for _, f := range configFiles {
		items = append(items, ListItem(f.Name()))
	}

	return items
}
