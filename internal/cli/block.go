package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"gatekeep/internal/config"
	"gatekeep/internal/country"
	"gatekeep/internal/shim"
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
	if args[0] == "gatekeep" {
		return fmt.Errorf("cannot block gatekeep itself")
	}
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
	if err := shim.Add(args[0]); err != nil {
		return err
	}

	rules, err := config.Load()
	if err != nil {
		return err
	}
	config.Add(rules, args[0], codes)
	if err := config.Save(rules); err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%s blocked in: %s\n", args[0], strings.Join(codes, ", "))
	return nil
}
