package utils

import (
	"os"
)

func GetHomeDir() (string, error) {
	return os.UserHomeDir()
}
