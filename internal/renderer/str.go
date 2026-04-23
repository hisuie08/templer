package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	var w = os.Stdout
	if opt.OutArg != "" {

		f, err := createFile(outPath)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}
	if opt.OutArg == "" {
		if !strings.HasSuffix(tmplStr, "\n") {
			tmplStr = fmt.Sprintln(tmplStr)
		}
	}
	return render("out", tmplStr, data, w)
}
