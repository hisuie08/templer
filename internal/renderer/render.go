package renderer

import (
	"io"
	"os"
	"path/filepath"
	"templer/internal/funcs"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func render(name, input string, data map[string]any, w io.Writer) error {

	t, err := newTmpl(name).Parse(input)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

func newTmpl(name string) *template.Template {
	return template.New(name).Funcs(sprig.FuncMap()).Funcs(funcs.Funcs())
}

func createFile(path string) (*os.File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
		return nil, err
	}
	return os.Create(abs)
}
