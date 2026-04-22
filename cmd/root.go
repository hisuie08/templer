package cmd

import (
	"fmt"
	"io"
	"os"
	"templer/internal/option"
	"templer/internal/process"

	"github.com/spf13/cobra"
)

var Version = "dev"
var opt option.Option
var rootCmd = &cobra.Command{
	Args: cobra.MaximumNArgs(1),
	Use:  "templer <template>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			opt.Template.Value = args[0]
		} else {
			b, err := readStdin()
			if err != nil {
				return err
			}
			if len(b) == 0 {
				return cmd.Root().Help()
			}
			opt.Template.Value = string(b)
		}
		p := process.New(opt)
		return p.Run()
	},
}

func readStdin() ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	// file or pipe
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return io.ReadAll(os.Stdin)
	}
	return nil, nil
}
func Execute() {
	rootCmd.Version = Version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&opt.Template.AsLiteral, "literal", false, "ensure template as string")
	rootCmd.Flags().StringArrayVarP(&opt.DataArgs, "data", "d", []string{}, "data file or string")
	rootCmd.Flags().StringVar(&opt.DataFormat, "data-format", "", "json|yaml")
	rootCmd.Flags().StringVarP(&opt.Template.Suffix, "suffix", "s", ".tmpl", "template file suffix")
	rootCmd.Flags().StringVarP(&opt.OutArg, "out", "o", "", "output path\nuse <template> path with the suffix trimmed when no arguments")
	rootCmd.Flags().StringArrayVar(&opt.SetValues, "set", nil, "Add K=V format entries file|string")
	rootCmd.Flags().BoolVar(&opt.LoadEnv, "env", true, "load env")
	rootCmd.Flag("out").NoOptDefVal = option.OutSibling
}
