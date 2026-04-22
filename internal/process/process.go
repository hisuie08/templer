package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	if !strings.HasPrefix(p.opt.Template.Suffix, ".") {
		p.opt.Template.Suffix = fmt.Sprintf(".%s", p.opt.Template.Suffix)
	}
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
