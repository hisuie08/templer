package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/option"
)

func RenderFile(opt option.Option, data map[string]any) error {
	outPath := func() string {
		filename := strings.TrimSuffix(
			filepath.Base(opt.Template.Value), opt.Template.Suffix)
		if opt.OutArg == option.OutSibling {
			return filepath.Join(filepath.Dir(opt.Template.Value), filename)
		}
		if i, err := os.Stat(opt.OutArg); err == nil {
			if i.IsDir() {
				return filepath.Join(opt.OutArg, filename)
			}
		}
		return opt.OutArg
	}()

	b, err := os.ReadFile(opt.Template.Value)
	if err != nil {
		fmt.Errorf(
			"cannot read path: %s: %w\nhint: use --literal to treat it as a string",
			opt.Template.Value, err,
		)
	}
	tmplStr := string(b)

	w, err := outWriter(outPath, opt.OutArg)
	if err != nil {
		return err
	}
	defer w.Close()

	return render("", fixForOut(tmplStr, opt.OutArg), data, w)
}
