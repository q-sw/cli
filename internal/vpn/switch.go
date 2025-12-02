package vpn

import (
	"fmt"
	"os"
	"strings"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func Connect(connectionName string) {

	var choice string

	if connectionName == "" {
		configs := utils.FetchFiles("vpnConfigPath")
		if len(configs) == 0 {
			fmt.Println("No VPN configurations found in", viper.GetString("vpnConfigPath"))
			return
		}
		choice = utils.List(configs)
	} else {
		choice = connectionName
	}

	if choice == "" {
		// User aborted the selection
		return
	}

	// Ensure the choice has the .conf suffix if provided via flag
	if connectionName != "" && !strings.HasSuffix(choice, ".conf") {
		choice += ".conf"
	}

	activeVpn, err := checkConnection()
	if err != nil {
		fmt.Println("Error checking VPN connection:", err)
		return
	}

	if activeVpn != "" {
		choiceNameWithoutExt := strings.TrimSuffix(choice, ".conf")
		if activeVpn == choiceNameWithoutExt {
			fmt.Printf("Already connected to %s.\n", choice)
			return
		}

		err := shutdown(activeVpn)
		if err != nil {
			fmt.Println("Error disconnecting from VPN:", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Connecting to %s...\n", choice)
	err = utils.ExecV("wg-quick", "up", viper.GetString("vpnConfigPath")+"/"+choice)
	if err != nil {
		fmt.Println("Error connecting to VPN:", err)
		return
	}
	fmt.Println("Successfully connected to", choice)
}
