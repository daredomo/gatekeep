package shim

import (
	"fmt"
	"slices"

	"gatekeep/internal/api"
	"gatekeep/internal/config"
)

func checkCountry(name string, args []string) error {
	rules, err := config.Load()
	if err != nil {
		return err
	}

	codes, ok := rules[name]
	if !ok {
		return nil
	}

	country, err := api.Country()
	if err != nil {
		return fmt.Errorf("cannot determine your country: %w", err)
	}
	if slices.Contains(codes, country) {
		return fmt.Errorf("%s is blocked in %s", name, country)
	}
	return nil
}
