package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"templer/internal/engine"
	"templer/internal/input"

	"github.com/spf13/cobra"
)

var (
	tmplArg     string
	tmplDir     string
	inputArg    string
	inputFormat string
	outArg      string
	setValues   []string
)
var rootCmd = &cobra.Command{
	Use: "templer",
	RunE: func(cmd *cobra.Command, args []string) error {

		data, err := input.Load(inputArg, inputFormat, setValues)
		if err != nil {
			return err
		}

		if tmplDir != "" {
			return engine.RenderDir(tmplDir, outArg, data)
		}
		if tmplArg == "" {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			tmplArg = filepath.Join(wd, "template.tmpl")
		}
		return engine.RenderOne(tmplArg, outArg, data)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVar(&tmplArg, "tmpl", "", "template file or string")
	rootCmd.Flags().StringVar(&tmplDir, "tmpl-dir", "", "template directory")

	rootCmd.Flags().StringVar(&inputArg, "input", "", "input file or string")
	rootCmd.Flags().StringVar(&inputFormat, "input-format", "", "json|yaml")

	rootCmd.Flags().StringVar(&outArg, "out", "", "output file or directory")

	// rootCmd.Flags().StringArrayVar(&setValues, "set", nil, "set values key=value")

}
