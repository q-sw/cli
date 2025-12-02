/*
Copyright © 2025 qsw
*/
package cmd

import (
	"log"
	"os"

	"github.com/q-sw/cli/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "QSW cli to be fast and productive",
	Long: `CLI to be fast and productive in my 
	SRE/DevOps work. Create new project, switch between config ...`,
}

func Execute() {
	err := cliCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cliInitConfig = &cobra.Command{
	Use:   "init",
	Short: "Init CLI config file",
	Run: func(cmd *cobra.Command, args []string) {
		cli.InitConfig(cliConfigFilePath)
	},
}

var cliConfigFilePath string

func init() {
	cobra.OnInitialize(initConfig)
	cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cliCmd.AddCommand(cliInitConfig)
	cliInitConfig.Flags().StringVarP(&cliConfigFilePath, "path", "p", "", "Config file parth defautl: $HOME/.config")

}

func initConfig() {

	home, _ := os.UserHomeDir()
	viper.AddConfigPath(home + "/.config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("cliconfig")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		log.Println("error to read config file")
	}
}
