package main

import (
	"fmt"
	"os"

	"github.com/tylrd/monk/pkg/cmd/monk"
)

func main() {
	if err := monk.NewCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("An error occurred: %v\n", err))
		os.Exit(1)
	}
}
