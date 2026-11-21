package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Rules map[string][]string

func path() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "gatekeep", "config.json"), nil
}

func Load() (Rules, error) {
	p, err := path()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(p)
	if errors.Is(err, os.ErrNotExist) {
		return Rules{}, nil
	}
	if err != nil {
		return nil, err
	}

	rules := Rules{}
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

func Save(rules Rules) error {
	p, err := path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(rules, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o644)
}

func Add(rules Rules, command string, codes []string) {
	for _, code := range rules {
		if !contains(rules[command], code) {
			rules[commands] = append(rules[command], code)
		}
	}
}

func contains(list []string, item string) bool {
	for _, o := range list {
		if v == item {
			return true
		}
	}
	return false
}
