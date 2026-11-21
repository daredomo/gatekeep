package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"gatekeep/internal/shim"
)

func newInstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Set up the shims directory and add it to your PATH",
		Args:  cobra.NoArgs,
		RunE:  runInstall,
	}
}

func runInstall(cmd *cobra.Command, _ []string) error {
	msg, err := shim.Setup()
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), msg)
	return nil
}
