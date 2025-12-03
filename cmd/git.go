/*
Copyright © 2025 q-sw
*/
package cmd

import (
	"log"

	"github.com/q-sw/cli/internal/git"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "work with git",
	Long:  `Work with git, switch config, init repo, get all repository info`,
}

var gitSwitch = &cobra.Command{
	Use:   "switch-config",
	Short: "Switch config",
	Long: `Switch your git config based on the CLI config,
	you must create gitConfig map with the key = config Name and the
	value = config file path`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := git.SwitchConfig(configName); err != nil {
			log.Fatalf("could not switch git config: %v", err)
		}
	},
}

var gitGetCurrentConfig = &cobra.Command{
	Use:   "get-config",
	Short: "get current config",

	Run: func(cmd *cobra.Command, args []string) {
		git.GetCurrentConfig()
	},
}

var gitGetStatus = &cobra.Command{
	Use:   "status",
	Short: "get git status of all repo",

	Run: func(cmd *cobra.Command, args []string) {
		git.GetDevStatus(statusVerbose, showBranch, showAllBranches)
	},
}

var configName string
var statusVerbose bool
var showChange bool
var showAllBranches bool
var showBranch bool

func init() {
	cliCmd.AddCommand(gitCmd)

	gitCmd.AddCommand(gitSwitch)
	gitSwitch.Flags().StringVarP(&configName, "name", "n", "", "-n or --name [config name]")

	gitCmd.AddCommand(gitGetCurrentConfig)

	gitCmd.AddCommand(gitGetStatus)
	gitGetStatus.Flags().BoolVarP(&statusVerbose, "verbose", "v", false, "[Global] Show details about repository status")
	gitGetStatus.Flags().BoolVarP(&showChange, "show-change", "c", false, "[Global] Show files changed")
	gitGetStatus.Flags().BoolVarP(&showBranch, "show-branch", "b", false, "[Global] Show actual branch")
	gitGetStatus.Flags().BoolVarP(&showAllBranches, "show-all-branches", "", false, "[Global] Show local and remote branches")
}
