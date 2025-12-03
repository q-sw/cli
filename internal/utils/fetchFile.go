package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/viper"
)

func FetchFiles(configPath string) ([]list.Item, error) {
	c := viper.GetString(configPath)

	if c == "" {
		return nil, fmt.Errorf("%s is not set in your config file", configPath)
	}

	configFiles, err := os.ReadDir(c)
	if err != nil {
		return nil, fmt.Errorf("error reading directory '%s': %w", c, err)
	}

	var items []list.Item
	for _, f := range configFiles {
		items = append(items, ListItem(f.Name()))
	}

	return items, nil
}
