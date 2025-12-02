package utils

import (
	"log"
	"os"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("error to get home user path")
		os.Exit(1)
	}
	return homeDir
}
