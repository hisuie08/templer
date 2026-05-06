package funcs

import (
	"maps"

	"github.com/Masterminds/sprig/v3"
)

func Funcs() map[string]any {
	funcMap :=sprig.FuncMap()
	custom:= map[string]any{
		"readFile": readFile,
		"Ternary":  ternary,
	}
	maps.Copy(funcMap,custom)
	return funcMap
}
