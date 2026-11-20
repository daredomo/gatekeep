package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "beta"

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the gatekeep version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "gatekeep", Version)
		},
	}
}
