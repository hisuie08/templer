package funcs

import (
	"maps"
	"templer/internal/option"
	"time"

	"github.com/Masterminds/sprig/v3"
)

type TemplerFunc struct {
	opt          option.Option
	funcMap      map[string]any
	shellTimeout time.Duration
}

func New(o option.Option) *TemplerFunc {
	funcMap := sprig.FuncMap()
	if !o.AllowEnv {
		delete(funcMap, "env")
		delete(funcMap, "expandenv")
	}
	t := &TemplerFunc{opt: o, shellTimeout: 3 * time.Second}
	maps.Copy(funcMap, map[string]any{
		"ReadFile": t.readFile,
		"Cwd":      t.cwd,
		"Exec":     t.execShell,
	})
	t.funcMap = funcMap
	return t
}

func (f *TemplerFunc) SetTimeout(t time.Duration) {
	f.shellTimeout = t
}

func (f *TemplerFunc) Funcs() map[string]any {
	return f.funcMap
}
