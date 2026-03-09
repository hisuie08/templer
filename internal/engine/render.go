package engine

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"templer/internal/files"

	"github.com/Masterminds/sprig/v3"
)

func RenderOne(tmplArg, out string, data map[string]any) error {

	tmplStr := loadTemplate(tmplArg)
	t := template.New("templer").
		Funcs(sprig.FuncMap()).
		Funcs(files.Funcs()) // readFile
	t, err := t.Parse(tmplStr)
	if err != nil {
		return err
	}

	var w = os.Stdout
	if out != "" {
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}

	return t.Execute(w, data)
}

func RenderDir(tmplDir, outDir string, data map[string]any) error {

	if outDir == "" {
		outDir = "out"
	}

	return filepath.WalkDir(tmplDir, func(path string, d os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		rel, err := filepath.Rel(tmplDir, path)
		if err != nil {
			return err
		}

		outPath := filepath.Join(outDir, strings.TrimSuffix(rel, ".tmpl"))

		err = os.MkdirAll(filepath.Dir(outPath), 0755)
		if err != nil {
			return err
		}

		tmplBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		t := template.New(rel).
			Funcs(sprig.FuncMap()).
			Funcs(files.Funcs())

		t, err = t.Parse(string(tmplBytes))
		if err != nil {
			return err
		}

		f, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer f.Close()

		return t.Execute(f, data)
	})
}
