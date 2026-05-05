package renderer

import (
	"os"
	"path/filepath"
	"templer/internal/context"
	"templer/internal/option"
)

type literalRenderer struct {
	ctx  context.Context
	opt  option.Option
	data map[string]any
}

func (l *literalRenderer) fixPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if l.opt.OutArg == option.OutSibling {
		return filepath.Join(wd, "out")
	}
	return filepath.Join(wd, l.opt.OutArg)
}

func (l *literalRenderer) Render() error {
	outPath := l.fixPath()
	tmplStr := l.opt.Template.Value
	return render("out", tmplStr, outPath, l.data, l.ctx.Out)
}
