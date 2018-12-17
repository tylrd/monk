package main

import (
	"fmt"
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("An error occurred: %v\n", err))
		os.Exit(1)
	}
}
