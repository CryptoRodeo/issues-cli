package main

import (
	"fmt"
	"os"

	"github.com/CryptoRodeo/issues-cli/cmd"
	"github.com/CryptoRodeo/issues-cli/pkg/config"
)

func main() {
	// Initialize configuration
	if err := config.InitConfig(); err != nil {
		fmt.Println("Error initializing config:", err)
		os.Exit(1)
	}

	// Execute the root command
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
