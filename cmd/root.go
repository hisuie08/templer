package cmd

import (
	"fmt"
	"io"
	"os"
	"templer/internal/context"
	"templer/internal/option"
	"templer/internal/output"
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
		var context = context.Context{
			Out:  output.OutController(cmd.OutOrStdout()),
			Log:  cmd.ErrOrStderr(),
			Err:  cmd.ErrOrStderr(),
		}
		p := &process.Process{Ctx: context, Opt: opt}
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

const data_help = `YAML (or JSON) format data <file|glob|str>:<value>
file:<value> -> force value as file to read
glob:<value> -> force value as glob pattern
str:<value> -> force value as string literal`

// TODO: templateもsuffixじゃなくてglobにしてもいいかも(検討中)
func init() {
	rootCmd.Flags().BoolVar(&opt.Template.AsLiteral, "literal", false, "ensure template as string")
	rootCmd.Flags().StringArrayVarP(&opt.DataArgs, "data", "d", []string{}, data_help)
	rootCmd.Flags().StringVar(&opt.DataFormat, "data-format", "", "force data format json|yaml\n In most cases, you don't need to specify it.")
	rootCmd.Flags().StringVar(&opt.Template.Suffix, "suffix", ".tmpl", "template file suffix")
	rootCmd.Flags().StringVarP(&opt.OutArg, "outdir", "o", "", "write output files under the specified directory")
	rootCmd.Flags().BoolVarP(&opt.OutDefault, "out", "O", false, "write output beside the template file, removing the template suffix")
	rootCmd.Flags().StringArrayVar(&opt.SetValues, "set", nil, "K=V data entries\nALWAYS overwrite duplicate entries in --data. file|string")
	rootCmd.Flags().BoolVar(&opt.IgnoreEnv, "no-env", false, "disable automatic loading of environment variables")
	rootCmd.Flags().BoolVar(&opt.AllowShellExecution,"allow-shell-execution",false,"enable function shell command execution")
}
