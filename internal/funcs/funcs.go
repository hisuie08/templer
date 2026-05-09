package funcs

import (
	"maps"
	"templer/internal/option"

	"github.com/Masterminds/sprig/v3"
)

type TemplerFunc struct {
	opt     option.Option
	funcMap map[string]any
}

func New(o option.Option) *TemplerFunc {
	funcMap := sprig.FuncMap()
	t := &TemplerFunc{opt: o}
	maps.Copy(funcMap, map[string]any{
		"ReadFile": t.readFile,
		"Cwd":      t.cwd,
		"Exec":     t.execShell,
	})
	t.funcMap = funcMap
	return t
}

func (f *TemplerFunc) Load(fm map[string]any) *TemplerFunc {
	maps.Copy(f.funcMap, fm)
	return f
}

func (f *TemplerFunc) Funcs() map[string]any {
	return f.funcMap
}
