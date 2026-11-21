package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "gatekeep",
		Short:        "Blocks terminal commands based on rules",
		SilenceUsage: true,
	}
}

func Execute() {
	root := newRootCmd()

	root.AddCommand(
		newInstallCmd(),
		newVersionCmd(),
		newBlockCmd(),
	)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
