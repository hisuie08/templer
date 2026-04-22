package renderer

import (
	"os"
	"path/filepath"
	"strings"

	"templer/internal/option"
)

func RenderDir(opt option.Option, data map[string]any) error {
	outDir := func() string {
		if opt.OutArg == "" {
			return opt.Template.Value
		}
		return opt.OutArg
	}()

	return filepath.WalkDir(opt.Template.Value, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, opt.Template.Suffix) {
			return nil
		}
		rel, err := filepath.Rel(opt.Template.Value, path)
		if err != nil {
			return err
		}

		outPath := filepath.Join(outDir, strings.TrimSuffix(rel, opt.Template.Suffix))

		f, err := createFile(outPath)
		defer f.Close()
		if err != nil {
			return err
		}

		tmplBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		t, err := newTmpl(rel).Parse(string(tmplBytes))
		if err != nil {
			return err
		}

		return t.Execute(f, data)
	})
}
