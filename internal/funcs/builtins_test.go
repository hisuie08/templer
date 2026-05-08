package funcs

import (
	"os"
	"path/filepath"
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
		name string // description of this test case
		// Named input parameters for target function.
		cmd     string
		args    []string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "pwd", cmd: "pwd", args: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := (&TemplerFunc{}).execShell(tt.cmd, tt.args...)
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
