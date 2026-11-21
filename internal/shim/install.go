package shim

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func shimsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "gatekeep", "shims"), nil
}

func Setup() (string, error) {
	dir, err := shimsDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	if filepath.Base(os.Getenv("SHELL")) != "fish" {
		return fmt.Sprintf("created %s\nadd it to the start of your PATH to finish the setup", dir), nil
	}

	cfg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	confDir := filepath.Join(cfg, "fish", "conf.d")
	if err := os.MkdirAll(confDir, 0o755); err != nil {
		return "", err
	}

	file := filepath.Join(confDir, "gatekeep.fish")
	line := fmt.Sprintf("fish_add_path -m %q\n", dir)
	if err := os.WriteFile(file, []byte(line), 0o644); err != nil {
		return "", err
	}

	return fmt.Sprintf("created %s\nwrote %s\nopen a new terminal to apply it", dir, file), nil
}

func Add(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("command %q not found", name)
	}

	dir, err := shimsDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	self, err := os.Executable()
	if err != nil {
		return err
	}

	err = os.Symlink(self, filepath.Join(dir, name))
	if errors.Is(err, fs.ErrExist) {
		return nil
	}
	return err
}
