package renderer

import (
	"os"
	"path/filepath"
	"templer/internal/files"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func newTmpl(name string) *template.Template {
	return template.New(name).Funcs(sprig.FuncMap()).Funcs(files.Funcs())
}

func createFile(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	return os.Create(path)
}
