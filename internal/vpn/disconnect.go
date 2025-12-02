package vpn

import (
	"fmt"
	"os"
)

func Disconnect() {
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
		fmt.Println("Successfully disconnected from", activeVpn)
	}

	fmt.Println("No connection active detected")

}
