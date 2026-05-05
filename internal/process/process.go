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
)

type Process struct {
	Ctx context.Context
	Opt option.Option
}

func (p *Process) Run() error {
	if !strings.HasPrefix(p.Opt.Template.Suffix, ".") {
		p.Opt.Template.Suffix = fmt.Sprintf(".%s", p.Opt.Template.Suffix)
	}
	data := parser.New(p.Opt).Parse()
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
