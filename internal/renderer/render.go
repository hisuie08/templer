package renderer

import (
	"bytes"
	"fmt"
	"strings"
	"templer/internal/context"
	"templer/internal/funcs"
	"templer/internal/option"
	"templer/internal/output"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func render(name, input string, outPath string,
	data map[string]any, out output.Output) error {
	if out.IsStd() {
		input = fixForOut(input)
	}
	t, err := newTmpl(name).Parse(input)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	if err := t.Execute(b, data); err != nil {
		return err
	}
	return out.WriteFile(outPath, b.Bytes())
}

func newTmpl(name string) *template.Template {
	return template.New(name).Funcs(sprig.FuncMap()).Funcs(funcs.Funcs())
}
func fixForOut(str string) string {
	result := str
	if !strings.HasSuffix(str, "\n") {
		result = fmt.Sprintln(str)
	}
	return result
}

type Renderer interface {
	Render() error
}

func Literal(ctx context.Context,
	opt option.Option, data map[string]any) Renderer {
	return &literalRenderer{ctx: ctx, opt: opt, data: data}
}

func File(ctx context.Context,
	opt option.Option, data map[string]any) Renderer {
	return &fileRenderer{ctx: ctx, opt: opt, data: data}
}
func Dir(ctx context.Context,
	opt option.Option, data map[string]any) Renderer {
	return &dirRenderer{ctx: ctx, opt: opt, data: data}
}
