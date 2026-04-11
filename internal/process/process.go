package process

import (
	"path/filepath"
	"templer/internal/data"
	"templer/internal/engine"
	"templer/internal/option"
)

type process struct {
	opt option.Option
}

func New(o option.Option) process {

	return process{opt: o}
}
func (p *process) Run() error {
	data, err := data.Load(p.opt.DataArgs, p.opt.DataFormat, p.opt.SetValues)
	if err != nil {
		return err
	}

	if p.opt.TemplateType() == "dir" {
		return engine.RenderDir(p.opt.TmplArg, p.opt.OutArg, data, p.opt.TmplSuffix)
	}
	return engine.RenderOne(p.opt.TmplArg, p.opt.OutArg, data)
}

func getOutPath(p string) (string, error) {
	return filepath.Abs(p)
}
