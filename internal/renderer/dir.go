package renderer

import (
	"fmt"
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

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		tmplStr := string(b)
		outPath, err := filepath.Abs(filepath.Join(outDir, strings.TrimSuffix(
			filepath.Base(path), opt.Template.Suffix)))
		if err != nil {
			return err
		}
		var w = os.Stdout
		if opt.OutArg != "" {
			w, err = createFile(outPath)
			if err != nil {
				return err
			}
			defer w.Close()
		}

		if opt.OutArg == "" {
			tmplStr = fmt.Sprintf("%s\n%s", outPath, tmplStr)
			if !strings.HasSuffix(tmplStr, "\n") {
				tmplStr = fmt.Sprintln(tmplStr)
			}
		}

		return render(path, string(tmplStr), data, w)
	})
}
