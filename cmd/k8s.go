/*
Copyright © 2025 q-sw
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/q-sw/cli/internal/k8s"
	"github.com/spf13/cobra"
)

var k8sContextName string

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "work with kubernetes",
	Long:  `Work with kubernetes, switch config, run complexe commands...`,
}

var k8sSwitchContext = &cobra.Command{
	Use:   "context",
	Short: "Switch kubernetes context",
	Run: func(cmd *cobra.Command, args []string) {
		chosenContext, err := k8s.SwitchContext(k8sContextName)
		if err != nil {
			log.Fatalf("could not switch k8s context: %v", err)
		}
		// If choice is empty, it means the user aborted the selection.
		// The internal function returns a nil error in this case.
		if chosenContext != "" {
			fmt.Printf("Switched to Kubernetes context: %s\n", chosenContext)
		}
	},
}

func init() {
	cliCmd.AddCommand(k8sCmd)
	k8sCmd.AddCommand(k8sSwitchContext)

	k8sSwitchContext.Flags().StringVarP(&k8sContextName, "name", "n", "", "Name of the kubernetes context to switch to")
}
