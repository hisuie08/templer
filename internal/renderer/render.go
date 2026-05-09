package renderer

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"templer/internal/context"
	"templer/internal/funcs"
	"templer/internal/option"
	"templer/internal/output"
	"text/template"
)

func render(name, input string, outPath string,
	data map[string]any, out output.Output, opt option.Option) error {
	t, err := newTmpl(name, opt).Parse(input)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	if err := t.Execute(b, data); err != nil {
		if errors.Is(err, funcs.ErrShellDisabled) {
			return funcs.ErrShellDisabled
		}
		return err
	}
	if out.IsStd() {
		b = bytes.NewBufferString(fixForOut(b.String()))
	}
	return out.WriteFile(outPath, b.Bytes())
}

func newTmpl(name string, o option.Option) *template.Template {
	return template.New(name).Funcs(funcs.New(o).Funcs())
}
func fixForOut(str string) string {
	result := str
	if !strings.HasSuffix(str, "\n") {
		result = fmt.Sprintln(str)
	}
	return result
}
func isValidDir(path string) (bool, error) {
	if i, err := os.Stat(path); err == nil {
		return i.IsDir(), nil
	} else {
		return false, err
	}
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
