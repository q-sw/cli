package vpn

import "fmt"

func Disconnect() (string, error) {
	activeVpn, err := checkConnection()
	if err != nil {
		return "", fmt.Errorf("error checking VPN connection: %w", err)
	}

	if activeVpn == "" {
		return "", nil
	}

	if err := shutdown(activeVpn); err != nil {
		return "", fmt.Errorf("error disconnecting from VPN: %w", err)
	}

	return activeVpn, nil
}
