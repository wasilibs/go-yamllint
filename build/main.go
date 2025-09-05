package main

import (
	"flag"

	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	// TODO: Investigate why yamllint ignore isn't working here.
	_ = flag.Lookup("skip").Value.Set("lint-yaml,format-yaml")
	tasks.Define(tasks.Params{
		LibraryName: "yamllint",
		LibraryRepo: "adrienverge/yamllint",
		GoReleaser:  true,
	})
	boot.Main()
}
