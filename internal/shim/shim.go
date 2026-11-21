package shim

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func Run(args []string) int {
	name := filepath.Base(args[0])

	if err := runMiddlewares(name, args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "gatekeep:", err)
		return 1
	}

	path, err := findReal(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "gatekeep:", err)
		return 127
	}

	err = syscall.Exec(path, args, os.Environ())
	fmt.Fprintln(os.Stderr, "gatekeep:", err)
	return 126
}

func findReal(name string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	shimsDir := filepath.Join(home, ".local", "share", "gatekeep", "shims")

	var dirs []string
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if filepath.Clean(dir) != shimsDir {
			dirs = append(dirs, dir)
		}
	}

	os.Setenv("PATH", strings.Join(dirs, string(os.PathListSeparator)))
	return exec.LookPath(name)
}
