package main

import (
	"os"

	"github.com/ickim/examplebeat/cmd"

	_ "github.com/ickim/examplebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
