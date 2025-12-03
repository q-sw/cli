package vpn

import (
	"path/filepath"
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
	vpnPath := filepath.Join(viper.GetString("vpnConfigPath"), connectName+".conf")
	err := utils.ExecV("wg-quick", "down", vpnPath)
	if err != nil {
		return err
	}
	return nil
}
