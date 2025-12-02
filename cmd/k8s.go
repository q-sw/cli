/*
Copyright © 2025 q-sw
*/
package cmd

import (
	"github.com/q-sw/cli/internal/k8s"
	"github.com/spf13/cobra"
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "work with kubernetes",
	Long:  `Work with kubernetes, switch config, run complexe commands...`,
}

var k8sSwitchContext = &cobra.Command{
	Use:   "context",
	Short: "Switch kubernetes context",
	Run: func(cmd *cobra.Command, args []string) {
		k8s.SwitchContext()
	},
}

func init() {
	cliCmd.AddCommand(k8sCmd)
	k8sCmd.AddCommand(k8sSwitchContext)

}
