package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/context"
	"templer/internal/option"
)

func (f *fileRenderer) fixPath() string {
	filename := strings.TrimSuffix(
		filepath.Base(f.opt.Template.Value), f.opt.Template.Suffix)
	if f.opt.OutArg == option.OutSibling {
		return filepath.Join(filepath.Dir(f.opt.Template.Value), filename)
	}
	if i, err := os.Stat(f.opt.OutArg); err == nil {
		if i.IsDir() {
			return filepath.Join(f.opt.OutArg, filename)
		}
	}
	return f.opt.OutArg
}

type fileRenderer struct {
	ctx  context.Context
	opt  option.Option
	data map[string]any
}

func (f *fileRenderer) Render() error {
	outPath := f.fixPath()
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
