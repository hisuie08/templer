package option

import (
	"os"
)

type Option struct {
	TmplArg    string
	TmplType   string
	TmplSuffix string
	DataArgs   []string
	DataFormat string
	OutArg     string
	SetValues  []string
	LoadEnv    bool
}

func (o *Option) TemplateType() string {
	if o.TmplType == "dir" || o.TmplType == "file" || o.TmplType == "string" {
		return o.TmplType
	}
	if fi, err := os.Stat(o.TmplArg); err == nil {
		if fi.IsDir() {
			return "dir"
		}
		return "file"
	}
	return "string"
}
