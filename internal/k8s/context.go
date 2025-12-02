package k8s

import (
	"fmt"
	"os"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func SwitchContext() {
	c := viper.GetString("kubeConfigPath")
	configs := utils.FetchFiles("kubeConfigPath")
	choice := utils.List(configs)

	if choice != "" {
		content, err := os.ReadFile(c + "/" + choice)
		if err != nil {
			fmt.Printf("Error reading file '%s': %v", choice, err)
			os.Exit(1)
		}
		home, _ := os.UserHomeDir()
		os.WriteFile(home+"/.kube/config", content, 0644)
		fmt.Println(string(content))
	}
}
