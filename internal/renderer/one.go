package renderer

import (
	"fmt"
	"os"
	"strings"
	"templer/internal/option"
)

func resolveTpl(opt option.Option) string {
	if opt.Template.AsLiteral {
		return opt.Template.Value
	}
	if b, err := os.ReadFile(opt.Template.Value); err == nil {
		return string(b)
	}
	return opt.Template.Value
}

func RenderOne(opt option.Option, data map[string]any) error {
	tmplStr := resolveTpl(opt)
	var w = os.Stdout
	if opt.OutArg != "" {
		f, err := os.Create(opt.OutArg)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	} else {
		if !strings.HasSuffix(tmplStr, "\n") {
			tmplStr = fmt.Sprintln(tmplStr)
		}
	}
	t := newTmpl("main")
	t, err := t.Parse(tmplStr)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
