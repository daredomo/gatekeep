package cli

import (
	"github.com/spf13/cobra"
)

func newBlockCmd() *cobra.Command {
	blockCmd := &cobra.Command{
		Use:   "block <command>",
		Short: "<write this>",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	blockCmd.Flags().StringSlice("in", nil, "country codes, e.g. MX,CO")
	blockCmd.MarkFlagRequired("in")

	return blockCmd
}
