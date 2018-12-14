package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Short: "Print the version of the monk client",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", "0.1.0")
		},
	}
	return cmd
}