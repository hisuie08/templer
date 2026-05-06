package funcs

import (
	"maps"

	"github.com/Masterminds/sprig/v3"
)

var builtInFuncs map[string]any = map[string]any{
	"readFile": readFile,
	"cwd":      cwd,
}

type TemplerFunc struct {
	funcMap map[string]any
}

func Funcs() map[string]any {
	funcMap := sprig.FuncMap()
	maps.Copy(funcMap, builtInFuncs)
	return funcMap
}

// TODO: カスタム関数動的読み込み未実装
func New() *TemplerFunc {
	funcMap := sprig.FuncMap()
	maps.Copy(funcMap, builtInFuncs)
	return &TemplerFunc{funcMap: funcMap}
}

func (f *TemplerFunc) Load(fm map[string]any) *TemplerFunc {
	maps.Copy(f.funcMap, fm)
	return f
}

func (f *TemplerFunc) Build() map[string]any {
	return f.funcMap
}
