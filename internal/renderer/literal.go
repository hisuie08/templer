package renderer

import (
	"fmt"
	"path/filepath"
	"templer/internal/context"
	"templer/internal/option"
)

type literalRenderer struct {
	ctx  context.Context
	opt  option.Option
	data map[string]any
}

func (l *literalRenderer) fixOut() (string,error) {
	if l.opt.OutArg !=""{
		dir,err:=isValidDir(filepath.Dir(l.opt.OutArg))
		if err!=nil{
			return "",err
		}
		if dir{
			return "",fmt.Errorf("invalid path: %s is a directory",l.opt.OutArg)
		}
		return filepath.Abs(l.opt.OutArg)
	}
	l.ctx.Out.AsStd()
	return "",nil
}

func (l *literalRenderer) Render() error {
	var outPath = ""
	if l.opt.OutArg != "" {
		outPath = filepath.Join(l.ctx.Root, l.opt.OutArg)
	} else {
		l.ctx.Out.AsStd()
	}

	tmplStr := l.opt.Template.Value
	return render("out", tmplStr, outPath, l.data, l.ctx.Out)
}
