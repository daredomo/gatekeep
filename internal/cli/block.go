package cli

import (
	"fmt"
	"gatekeep/internal/country"
	"strings"

	"github.com/spf13/cobra"
)

func newBlockCmd() *cobra.Command {
	blockCmd := &cobra.Command{
		Use:   "block <command>",
		Short: "<write this>",
		Args:  cobra.ExactArgs(1),
		RunE:  runBlock,
	}
	blockCmd.Flags().StringSlice("in", nil, "country codes, e.g. MX,CO")
	blockCmd.MarkFlagRequired("in")

	return blockCmd
}

func runBlock(cmd *cobra.Command, args []string) error {
	countries, err := cmd.Flags().GetStringSlice("in")
	if err != nil {
		return err
	}

	var codes []string
	for _, c := range countries {
		code, err := country.NormalizeCode(c)
		if err != nil {
			return err
		}
		codes = append(codes, code)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%s blocked in: %s\n", args[0], strings.Join(codes, ", "))
	return nil
}
