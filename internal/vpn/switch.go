package vpn

import (
	"fmt"
	"os"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func Connect() {

	configs := utils.FetchFiles("vpnConfigPath")
	choice := utils.List(configs)

	if choice != "" {
		activeVpn, err := checkConnection()
		if err != nil {
			fmt.Println("Error checking VPN connection:", err)
			return
		}

		if activeVpn != "" {
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
}
