package main

import (
	"os"

	"github.com/Hayao0819/awesome-gomamayo/cmd"
)

func main() {
	if err := cmd.Root().Execute(); err != nil {
		os.Exit(1)
	}
}
