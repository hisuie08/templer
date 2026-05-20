package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/context"
	"templer/internal/option"
	"templer/internal/parser"
	"templer/internal/renderer"

	"gopkg.in/yaml.v3"
)

type Process struct {
	Ctx context.Context
	Opt option.Option
}

func (p *Process) Run() error {
	if !strings.HasPrefix(p.Opt.Template.Suffix, ".") {
		p.Opt.Template.Suffix = fmt.Sprintf(".%s", p.Opt.Template.Suffix)
	}
	data, err := parser.New(p.Opt, p.Ctx).Parse()
	if p.Opt.WithInspect {
		b, err := yaml.Marshal(data)
		if err != nil {
			return err
		}
		p.Ctx.Log.Write(append(b, []byte("\n")...))
	}
	if err != nil {
		return err
	}
	if p.Opt.Template.AsLiteral {
		return renderer.Literal(p.Ctx, p.Opt, data).Render()
	}
	fi, err := os.Stat(p.Opt.Template.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return renderer.Literal(p.Ctx, p.Opt, data).Render()
		}
		return err
	}
	if fi.IsDir() {
		return renderer.Dir(p.Ctx, p.Opt, data).Render()
	} else {
		return renderer.File(p.Ctx, p.Opt, data).Render()
	}
}

func getOutPath(p string) (string, error) {
	return filepath.Abs(p)
}
