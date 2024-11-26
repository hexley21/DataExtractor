package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/hexley21/data_extractor/cmd/ui/multi_select"
	"github.com/hexley21/data_extractor/pkg/extractor"
	"github.com/hexley21/data_extractor/pkg/serialization"
	"github.com/hexley21/data_extractor/pkg/util"
	"github.com/spf13/cobra"
)

var (
	ErrNoFileName = errors.New("no filename provided, shutting down")

	beautify bool
	rootCmd = &cobra.Command{
		Use:   "data-extractor [file]",
		Short: "A CLI for extracting specified fields from any file format",
		Long: `DataExtractor is a CLI app that can extract only provided fields from a file and output the new resulting file.
for now, only JSON extension is supported.`,
		Args: cobra.ExactArgs(1),
		RunE: RunE,
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&beautify, "beautify", "b", false, "Enable beautify mode")
}

func Run() error {
	return rootCmd.Execute()
}


func RunE(cmd *cobra.Command, args []string) error {
	return nil
}
