package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func SwitchContext(contextName string) (string, error) {
	configPath := viper.GetString("kubeConfigPath")
	var choice string

	if contextName == "" {
		configs, err := utils.FetchFiles("kubeConfigPath")
		if err != nil {
			return "", err
		}
		if len(configs) == 0 {
			return "", fmt.Errorf("no Kubernetes contexts found in %s", configPath)
		}
		choice = utils.List(configs)
	} else {
		choice = contextName
	}

	if choice == "" {
		// User aborted the selection from TUI
		return "", nil
	}

	fullPath := filepath.Join(configPath, choice)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("kubernetes context file '%s' not found", choice)
		}
		return "", fmt.Errorf("error reading file '%s': %w", choice, err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}

	kubeConfigPath := filepath.Join(home, ".kube", "config")
	err = os.WriteFile(kubeConfigPath, content, 0644)
	if err != nil {
		return "", fmt.Errorf("error writing kubeconfig file: %w", err)
	}

	return choice, nil
}
