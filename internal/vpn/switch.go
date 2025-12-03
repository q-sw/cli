package vpn

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/q-sw/cli/internal/utils"
	"github.com/spf13/viper"
)

func Connect(connectionName string) (string, error) {
	var choice string

	if connectionName == "" {
		configs, err := utils.FetchFiles("vpnConfigPath")
		if err != nil {
			return "", err
		}
		if len(configs) == 0 {
			return "", fmt.Errorf("no VPN configurations found in %s", viper.GetString("vpnConfigPath"))
		}
		choice = utils.List(configs)
	} else {
		choice = connectionName
	}

	if choice == "" {
		// User aborted the selection
		return "", nil
	}

	// Ensure the choice has the .conf suffix if provided via flag
	if connectionName != "" && !strings.HasSuffix(choice, ".conf") {
		choice += ".conf"
	}

	activeVpn, err := checkConnection()
	if err != nil {
		return "", fmt.Errorf("error checking VPN connection: %w", err)
	}

	if activeVpn != "" {
		choiceNameWithoutExt := strings.TrimSuffix(choice, ".conf")
		if activeVpn == choiceNameWithoutExt {
			return "", fmt.Errorf("already connected to %s", choice)
		}

		if err := shutdown(activeVpn); err != nil {
			return "", fmt.Errorf("error disconnecting from current VPN: %w", err)
		}
	}

	vpnPath := filepath.Join(viper.GetString("vpnConfigPath"), choice)
	if err := utils.ExecV("wg-quick", "up", vpnPath); err != nil {
		return "", fmt.Errorf("error connecting to VPN: %w", err)
	}
	return choice, nil
}
