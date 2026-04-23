package renderer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/funcs"
	"templer/internal/option"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func render(name, input string, data map[string]any, w io.WriteCloser) error {

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

type noopWriteCloser struct {
	io.Writer
}

func (noopWriteCloser) Close() error {
	// Override the method to prevent Stdout from being closed.
	return nil
}
func outWriter(outpath, outArg string) (io.WriteCloser, error) {
	if outArg != "" {
		return createFile(outpath)
	}

	return noopWriteCloser{os.Stdout}, nil
}

type renderer struct {
	opt     option.Option
	tmplStr string
}

func fixForOut(tmplStr, outArg string) string {
	result := tmplStr
	if outArg == "" {
		if !strings.HasSuffix(tmplStr, "\n") {
			result = fmt.Sprintln(tmplStr)
		}
	}
	return result
}
