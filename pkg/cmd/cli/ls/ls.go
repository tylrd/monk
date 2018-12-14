package ls

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tylrd/monk/pkg/cmd"
)

type ListOptions struct {
	ExcludeFolders bool
	Recursive      bool
}

func NewCommand() *cobra.Command {
	o := &ListOptions{}
	c := &cobra.Command{
		Use:   "ls",
		Short: "List collections for the active context",
		Run: func(c *cobra.Command, args []string) {
			cmd.CheckError(Run(o));
		},
	}
	c.Flags().BoolVarP(&o.Recursive, "recurse", "r", false, "Print all all collections recursively")
	c.Flags().BoolVarP(&o.ExcludeFolders, "exclude-folders", "e", false, "Exclude directories when printing out collections")
	return c
}

func Run(o *ListOptions) error {
	fmt.Println(o)
	return nil
}
