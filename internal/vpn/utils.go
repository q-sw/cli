package vpn

import (
	"fmt"
	"strings"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func checkConnection() (string, error) {
	out, err := utils.Exec("wg", "show", "interfaces")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

func shutdown(connectName string) error {
	fmt.Printf("Disconnecting from %s...\n", connectName)
	err := utils.ExecV("wg-quick", "down", viper.GetString("vpnConfigPath")+"/"+connectName+".conf")
	if err != nil {
		return err
	}
	return nil
}
