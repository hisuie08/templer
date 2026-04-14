package process

import (
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
	parser := parser.New(p.opt)
	data := parser.Parse()
	if p.opt.TemplateType() == "dir" {
		return renderer.RenderDir(p.opt.TmplArg, p.opt.OutArg, data, p.opt.TmplSuffix)
	}
	return renderer.RenderOne(p.opt.TmplArg, p.opt.OutArg, data)
}

func getOutPath(p string) (string, error) {
	return filepath.Abs(p)
}
