package funcs

import "github.com/expr-lang/expr"

func Call(name, fn string, args map[string]any) (any, error) {
	f, err := expr.Compile(fn)
	if err != nil {
		return nil, err
	}
	return expr.Run(f, args)
}
