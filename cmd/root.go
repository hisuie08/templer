package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"templer/internal/context"
	"templer/internal/funcs"
	"templer/internal/option"
	"templer/internal/output"
	"templer/internal/process"

	"github.com/spf13/cobra"
)

var Version = "dev"
var opt option.Option
var rootCmd = &cobra.Command{
	Args: cobra.MaximumNArgs(1),
	Use:  "templer <template-file|dir|string>",
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
			Out: output.OutController(cmd.OutOrStdout()),
			Log: cmd.ErrOrStderr(),
			Err: cmd.ErrOrStderr(),
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
		if errors.Is(err, funcs.ErrShellDisabled) ||
			errors.Is(err, funcs.ErrShellDisallowed) {
			rootCmd.PrintErrln(funcs.ShellExecWarning)
		}
		os.Exit(1)
	}
}

const data_help = `YAML/JSON data source
  auto-detects file, glob, or string
  prefixes:
    file:<v> force file
    glob:<v> force glob
    str:<v>  force string`
const out_help = `output destination
  - file path for single-file output
  - directory path for directory template output`
const out_default_help = `write output beside the template, removing the template suffix
only works when <template> is a file or directory.`

// TODO: templateもsuffixじゃなくてglobにしてもいいかも(検討中)
func init() {
	rootCmd.Flags().BoolVar(&opt.Template.AsLiteral, "literal", false, "force <template> as literal string")
	rootCmd.Flags().StringArrayVarP(&opt.DataArgs, "data", "d", []string{}, data_help)
	rootCmd.Flags().BoolVar(&opt.DataStrictJson, "strict-json", false, "parse --data values as strict JSON only")
	rootCmd.Flags().StringVar(&opt.Template.Suffix, "suffix", ".tmpl", "template file suffix")
	rootCmd.Flags().StringVarP(&opt.OutDir, "out", "o", "", out_help)
	rootCmd.Flags().BoolVarP(&opt.OutDefault, "out-default", "O", false, out_default_help)
	rootCmd.Flags().StringArrayVar(&opt.SetValues, "set", nil, "K=V data entries\nduplicate keys always override --data entries")
	rootCmd.Flags().BoolVar(&opt.AllowEnv, "env", false, "enable automatic loading of environment variables")
	rootCmd.Flags().BoolVar(&opt.AllowShellExecution, "allow-shell-execution", false, "enable function shell command execution")
	rootCmd.Flags().StringArrayVar(&opt.AllowedShell, "allow-command", []string{}, "specify commands to allow execution")
}
