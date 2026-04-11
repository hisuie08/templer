package cmd

import (
	"fmt"
	"os"

	"templer/internal/option"
	"templer/internal/process"

	"github.com/spf13/cobra"
)

var opt option.Option
var rootCmd = &cobra.Command{
	Args: cobra.ExactArgs(1),
	Use:  "templer <template> [OPTIONS]",
	RunE: func(cmd *cobra.Command, args []string) error {
		opt.TmplArg = args[0]
		p := process.New(opt)
		return p.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&opt.TmplType, "tmpl-type", "", "template type file|dir|string")
	rootCmd.Flags().StringArrayVarP(&opt.DataArgs, "data", "d", []string{}, "data file or string")
	rootCmd.Flags().StringVar(&opt.DataFormat, "data-format", "", "json|yaml")
	rootCmd.Flags().StringVar(&opt.TmplSuffix, "suffix", ".tmpl", "template file suffix")
	rootCmd.Flags().StringVar(&opt.OutArg, "out", "", "output file or directory")
	rootCmd.Flags().StringArrayVar(&opt.SetValues, "set", nil, "set values key=value")

}
