package process

import (
	"os"
	"path/filepath"
	"templer/internal/option"
	"templer/internal/parser"
	"templer/internal/renderer"
)

type process struct {
	opt option.Option
}

func New(o option.Option) process {
	return process{opt: o}
}
func (p *process) Run() error {
	data := parser.New(p.opt).Parse()
	if p.opt.Template.AsLiteral {
		return renderer.RenderStr(p.opt, data)
	}
	fi, err := os.Stat(p.opt.Template.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return renderer.RenderStr(p.opt, data)
		}
		return err
	}
	if fi.IsDir() {
		return renderer.RenderDir(p.opt, data)
	} else {
		return renderer.RenderFile(p.opt, data)
	}
}

func getOutPath(p string) (string, error) {
	return filepath.Abs(p)
}
