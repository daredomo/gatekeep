package main

import (
	"os"
	"path/filepath"

	"gatekeep/internal/cli"
	"gatekeep/internal/shim"
)

func main() {
	if filepath.Base(os.Args[0]) != "gatekeep" {
		os.Exit(shim.Run(os.Args))
	}
	cli.Execute()
}
