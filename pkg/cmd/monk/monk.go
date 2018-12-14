package monk

import (
	"github.com/spf13/cobra"
	"github.com/tylrd/monk/pkg/cmd/cli/create"
	"github.com/tylrd/monk/pkg/cmd/cli/ls"
	"github.com/tylrd/monk/pkg/cmd/version"
)

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "monk",
		Short: "Run and test HTTP collections",
		Long:  `Monk is a CLI for executing, managing, and testing HTTP collections.`,
	}
	c.AddCommand(
		version.NewCommand(),
		ls.NewCommand(),
		create.NewCommand(),
	)
	return c
}
