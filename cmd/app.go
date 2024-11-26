package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/hexley21/data_extractor/cmd/ui/multi_select"
	"github.com/hexley21/data_extractor/pkg/config"
	"github.com/hexley21/data_extractor/pkg/extractor"
	"github.com/hexley21/data_extractor/pkg/serialization"
	"github.com/hexley21/data_extractor/pkg/util"
	"github.com/spf13/cobra"
)

var (
	ErrNoFileName    = errors.New("no filename provided, shutting down")
	ErrInvalidIndent = errors.New("invalid indent, positive integer expected")

	cfg    *config.Config
	indent int
)

func Run(c *config.Config) error {
	cfg = c

	rootCmd := &cobra.Command{
		Use:   c.CliName + " [file]",
		Short: "A CLI for extracting specified fields from any file format",
		Long:  c.AppName + " is a CLI app that can extract only provided fields from a file and output the new resulting file.",
		Args:  cobra.ExactArgs(1),
		RunE:  RunE,
	}

	rootCmd.Flags().IntVarP(&indent, "indent", "i", 4, "Specify indent")

	return rootCmd.Execute()
}

func RunE(cmd *cobra.Command, args []string) error {
	fileArg := args[0]

	if fileArg == "" {
		return ErrNoFileName
	}
	if indent < 0 {
		return ErrInvalidIndent
	}

	fileExt := filepath.Ext(fileArg)

	processor, err := serialization.GetProcessor(fileExt, indent)
	if err != nil {
		return err
	}

	if !strings.Contains(fileArg, string(os.PathSeparator)) {
		fileArg, err = util.PrependCurrentDir(fileArg)
		if err != nil {
			return err
		}
	}

	content, err := util.ReadFromPath(fileArg)
	if err != nil {
		return err
	}

	var data interface{}
	err = processor.Deserialize(content, &data)
	if err != nil {
		return err
	}
	keys, viewKeys := extractor.New().Keys(data)

	selected, err := multi_select.DisplayChecklist(keys, viewKeys, cfg.Colors.MultiSelect[fileExt])
	if err != nil {
		return err
	}

	if selected == nil {
		return nil
	}

	data = extractor.New().Data(data, selected)

	res, err := processor.Serialize(data)
	if err != nil {
		return err
	}

	outputPath, err := util.PrependCurrentDir("extracted_" + filepath.Base(fileArg))
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, res, 0644)
}
