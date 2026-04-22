package renderer

import (
	"os"
	"path/filepath"
	"templer/internal/funcs"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func newTmpl(name string) *template.Template {
	return template.New(name).Funcs(sprig.FuncMap()).Funcs(funcs.Funcs())
}

func createFile(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	return os.Create(path)
}
