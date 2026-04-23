package renderer

import (
	"os"
	"path/filepath"
	"strings"

	"templer/internal/option"
)

func RenderDir(opt option.Option, data map[string]any) error {
	outDir := func() string {
		if opt.OutArg == option.OutSibling {
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
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		tmplStr := string(b)
		outPath := filepath.Join(outDir, strings.TrimSuffix(
			rel, opt.Template.Suffix))

		w, err := outWriter(outPath, opt.OutArg)
		if err != nil {
			return err
		}
		defer w.Close()
		return render(path, fixForOut(tmplStr, opt.OutArg), data, w)
	})
}
