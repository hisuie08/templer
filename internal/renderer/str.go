package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/option"
)

func resolveTpl(opt option.Option) string {
	if opt.Template.AsLiteral {
		return opt.Template.Value
	}
	b, err := os.ReadFile(opt.Template.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return opt.Template.Value
		}
	}
	if err == nil {
		return string(b)
	}
	return opt.Template.Value
}

func RenderStr(opt option.Option, data map[string]any) error {
	outDir := func() string {
		if opt.OutArg == option.OutSibling {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			return wd
		}
		return opt.OutArg
	}()
	tmplStr := opt.Template.Value
	var w = os.Stdout
	if opt.OutArg != "" {
		f, err := createFile(filepath.Join(outDir, "out"))
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
