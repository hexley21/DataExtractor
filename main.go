package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/hexley21/data_extractor/cmd"
	"github.com/hexley21/data_extractor/pkg/config"
)

//go:embed config.yml
var cfgFile []byte

func main() {
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := cmd.Run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
