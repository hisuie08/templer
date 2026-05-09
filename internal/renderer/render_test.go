package renderer

import (
	"os"
	"path/filepath"
	"templer/internal/option"
	"templer/internal/output"
	"testing"
)

func Test_fixForOut(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{name: "fix", str: "hello", want: "hello\n"},
		{name: "no fix", str: "hello\n", want: "hello\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fixForOut(tt.str)
			if got != tt.want {
				t.Fatalf("want: %s but got %s", tt.want, got)
			}
		})
	}
}

func Test_isValidDir(t *testing.T) {
	td := t.TempDir()
	nodir := filepath.Join(td, "notexist")
	file := filepath.Join(td, "file")
	if err := os.WriteFile(file, []byte("file"), 0644); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		path    string
		want    bool
		wantErr bool
	}{
		{name: "is dir", path: td, want: true, wantErr: false},
		{name: "is not dir", path: nodir, want: false, wantErr: true},
		{name: "is file", path: file, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := isValidDir(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("isValidDir() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("isValidDir() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isValidDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_render(t *testing.T) {
	tests := []struct {
		name    string
		t_name  string
		input   string
		outPath string
		data    map[string]any
		out     output.Output
		opt     option.Option
		wantErr bool
	}{{
		name: "test", t_name: "test",
		input:   "{{Cwd}}{{Exec \"ls\"}}{{upper .value}}",
		outPath: "/dev/null", data: map[string]any{"value": "test"},
		out: output.OutController(t.Output()),
		opt: option.Option{AllowShellExecution: true}, wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := render(tt.name, tt.input, tt.outPath, tt.data, tt.out, tt.opt)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("render() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("render() succeeded unexpectedly")
			}
		})
	}
}
