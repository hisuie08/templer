package cmd

import (
	"fmt"
	"os"

	"templer/internal/option"
	"templer/internal/process"

	"github.com/spf13/cobra"
)

var Version = "dev"
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
	rootCmd.Version = Version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&opt.TmplType, "tmpl-type", "t", "", "template type file|dir|string")
	rootCmd.Flags().StringArrayVarP(&opt.DataArgs, "data", "d", []string{}, "data file or string")
	rootCmd.Flags().StringVarP(&opt.DataFormat, "data-format", "f", "", "json|yaml")
	rootCmd.Flags().StringVarP(&opt.TmplSuffix, "suffix", "s", ".tmpl", "template file suffix")
	rootCmd.Flags().StringVarP(&opt.OutArg, "out", "o", "", "output file or directory")
	rootCmd.Flags().StringArrayVar(&opt.SetValues, "set", nil, "Add K=V format entries file|string")
	rootCmd.Flags().BoolVar(&opt.LoadEnv, "env", true, "load env")
}
