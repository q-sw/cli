package k8s

import (
	"fmt"
	"os"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func SwitchContext(contextName string) {
	c := viper.GetString("kubeConfigPath")
	var choice string

	if contextName == "" {
		configs := utils.FetchFiles("kubeConfigPath")
		if len(configs) == 0 {
			fmt.Println("No Kubernetes contexts found in", viper.GetString("kubeConfigPath"))
			return
		}
		choice = utils.List(configs)
	} else {
		choice = contextName
	}

	if choice == "" {
		// User aborted the selection
		return
	}

	fullPath := c + "/" + choice
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Printf("Kubernetes context file '%s' not found.\n", choice)
		return
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v", choice, err)
		os.Exit(1)
	}

	home, _ := os.UserHomeDir()
	err = os.WriteFile(home+"/.kube/config", content, 0644)
	if err != nil {
		fmt.Printf("Error writing Kubeconfig file: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Switched to Kubernetes context: %s\n", choice)
}
