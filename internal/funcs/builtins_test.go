package funcs

import (
	"os"
	"path/filepath"
	"templer/internal/option"
	"testing"
)

func Test_readFile(t *testing.T) {
	sp := filepath.Join(t.TempDir(), "exists.txt")
	content := "successful"
	if err := os.WriteFile(sp, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{name: "exists", path: sp, want: content, wantErr: false},
		{name: "not exists", path: "./fail.txt", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := (&TemplerFunc{}).readFile(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("readFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("readFile() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shell(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		allow   bool
		args    []string
		want    string
		wantErr bool
	}{

		{name: "disallow", cmd: "pwd", args: []string{},
			allow: false, wantErr: true},
		{name: "allow", cmd: "pwd", args: []string{},
			allow: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := &TemplerFunc{opt: option.Option{AllowShellExecution: tt.allow}}
			_, gotErr := tf.execShell(tt.cmd, tt.args...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("shell() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("shell() succeeded unexpectedly")
			}
		})
	}
}
