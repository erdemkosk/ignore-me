package main

import (
	"os"

	"github.com/erdemkosk/ignore-me/cmd"
	"github.com/fatih/color"
)

func main() {
	if err := cmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}
