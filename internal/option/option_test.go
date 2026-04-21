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
		name        string
		arg         string
		literal     bool
		wantDir     bool
		wantLiteral bool
	}{
		{name: "string", arg: "test", wantDir: false, wantLiteral: true},
		{name: "file", arg: filename, wantDir: false, wantLiteral: false},
		{name: "dir", arg: dirname, wantDir: true, wantLiteral: false},
		{name: "literal-file", arg: filename, literal: true, wantDir: false, wantLiteral: true},
		{name: "literal-dir", arg: dirname, literal: true, wantDir: false, wantLiteral: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o Option
			o.Template.AsLiteral = tt.literal || false
			o.Template.Value = tt.arg
			fi, err := os.Stat(tt.arg)
			if err != nil {
				if !tt.wantLiteral {
					t.Fatalf("expected literal: %v but not", tt.wantLiteral)
				}
			} else {
				if !tt.literal && fi.IsDir() != tt.wantDir {
					t.Errorf("expected dir: %v but got %v", fi.IsDir(), tt.wantDir)
				}
			}
		})
	}
}
