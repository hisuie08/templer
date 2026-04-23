package renderer

import (
	"os"
	"path/filepath"
	"templer/internal/option"
)

func RenderStr(opt option.Option, data map[string]any) error {
	outPath := func() string {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		if opt.OutArg == option.OutSibling {
			return filepath.Join(wd, "out")
		}
		return filepath.Join(wd, opt.OutArg)
	}()
	tmplStr := opt.Template.Value
	w, err := outWriter(outPath, opt.OutArg)
	if err != nil {
		return err
	}
	defer w.Close()

	return render("out", fixForOut(tmplStr, opt.OutArg), data, w)
}
