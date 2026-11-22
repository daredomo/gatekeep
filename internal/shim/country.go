package shim

import (
	"fmt"
	"slices"

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

	country := "MX" // extern api connection here
	if slices.Contains(codes, country) {
		return fmt.Errorf("%s is blocked in %s", name, country)
	}
	return nil
}
