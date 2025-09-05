package main

import (
	"os"

	"github.com/wasilibs/go-yamllint/internal/runner"
)

func main() {
	os.Exit(runner.Run("yamllint", os.Args[1:], os.Stdin, os.Stdout, os.Stderr, "."))
}
