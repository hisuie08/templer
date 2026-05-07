package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/context"
	"templer/internal/option"
)

func (f *fileRenderer) fixOut() (string, error) {
	filename := strings.TrimSuffix(
		filepath.Base(f.opt.Template.Value), f.opt.Template.Suffix)
	if f.opt.OutArg != "" { // --outdir / -o
		dir, err := isValidDir(f.opt.OutArg)
		if err != nil {
			return "", err
		}
		if dir {
			return filepath.Abs(filepath.Join(f.opt.OutArg, filename))
		} else {
			return filepath.Abs(f.opt.OutArg)
		}
	}
	if f.opt.OutDefault { // --out / -O
		return filepath.Abs(
			filepath.Join(filepath.Dir(f.opt.Template.Value), filename))
	}
	f.ctx.Out.AsStd()
	return "", nil
}

type fileRenderer struct {
	ctx  context.Context
	opt  option.Option
	data map[string]any
}

func (f *fileRenderer) Render() error {
	outPath, err := f.fixOut()
	if err != nil {
		return err
	}
	file := f.opt.Template.Value
	r, err := os.ReadFile(file)
	if err != nil {
		msg := fmt.Sprintf(
			"cannot read path: %s: %v\nhint: use --literal to treat it as a string\n",
			file, err,
		)
		f.ctx.Err.Write([]byte(msg))
	}
	tmplStr := string(r)
	return render(file, tmplStr, outPath, f.data, f.ctx.Out)
}
