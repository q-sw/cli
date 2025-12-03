/*
Copyright © 2025 q-sw
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/q-sw/cli/internal/vpn"
	"github.com/spf13/cobra"
)

var vpnName string

var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "manage VPN config",
}

var vpnConnect = &cobra.Command{
	Use:   "connect",
	Short: "Connect to VPN endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		connection, err := vpn.Connect(vpnName)
		if err != nil {
			log.Fatalf("could not connect to VPN: %v", err)
		}
		if connection != "" {
			fmt.Printf("Successfully connected to %s\n", connection)
		}
	},
}

var vpnDisconnect = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from VPN endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		disconnectedVpn, err := vpn.Disconnect()
		if err != nil {
			log.Fatalf("could not disconnect from VPN: %v", err)
		}
		if disconnectedVpn != "" {
			fmt.Printf("Successfully disconnected from %s\n", disconnectedVpn)
		} else {
			fmt.Println("No active VPN connection detected.")
		}
	},
}

func init() {
	cliCmd.AddCommand(vpnCmd)
	vpnCmd.AddCommand(vpnConnect)
	vpnCmd.AddCommand(vpnDisconnect)

	vpnConnect.Flags().StringVarP(&vpnName, "name", "n", "", "Name of the vpn to connect to")
}
