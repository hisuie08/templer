package option

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOption_TemplateType(t *testing.T) {
	dirname := t.TempDir()
	filename := filepath.Join(dirname, "test.tmpl")
	os.Mkdir(dirname, 0o666)
	os.Create(filename)
	tests := []struct {
		name  string
		arg   string
		ttype string
		want  string
	}{
		{name: "string", arg: "test", ttype: "string", want: "string"},
		{name: "file", arg: filename, ttype: "file", want: "file"},
		{name: "dir", arg: dirname, ttype: "dir", want: "dir"},
		{name: "u-string", arg: "test", want: "string"},
		{name: "u-file", arg: filename, want: "file"},
		{name: "u-dir", arg: dirname, want: "dir"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o Option
			o.TmplType = tt.ttype
			o.TmplArg = tt.arg
			got := o.TemplateType()
			if got != tt.want {
				t.Errorf("TemplateType() = %v, want %v", got, tt.want)
			}
		})
	}
}
