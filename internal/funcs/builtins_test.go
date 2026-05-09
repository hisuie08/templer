package funcs

import (
	"errors"
	"os"
	"path/filepath"
	"templer/internal/option"
	"testing"
	"time"
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

func Test_execShell(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		enable  bool
		allowed []string
		args    []string
		wantErr error
	}{

		{name: "disable-disallow", cmd: "pwd", enable: false, wantErr: ErrShellDisabled},
		{name: "disable-allow", cmd: "pwd", enable: false,
			allowed: []string{"pwd"}, wantErr: ErrShellDisabled},
		{name: "enable-allow", cmd: "pwd", enable: true,
			allowed: []string{"pwd"}, wantErr: nil},
		{name: "enable-disallow", cmd: "pwd", enable: true,
			allowed: []string{}, wantErr: ErrShellDisallowed},
		{name: "output-limit", cmd: "yes", allowed: []string{"yes"},
			enable: true, wantErr: ErrOutputLimitExceeded},
		{name: "timeout", cmd: "sleep", args: []string{"1"}, enable: true,
			allowed: []string{"sleep"}, wantErr: ErrShellTimeout},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := &TemplerFunc{opt: option.Option{
				AllowShellExecution: tt.enable, AllowedShell: tt.allowed},
				shellTimeout: 50 * time.Millisecond}
			_, gotErr := tf.execShell(tt.cmd, tt.args...)
			if tt.wantErr == nil {
				if gotErr != nil {
					t.Fatalf("unexpected error: %v", gotErr)
				}
				return
			}
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf(
					"expected error %v, got %v",
					tt.wantErr,
					gotErr,
				)
			}
		})
	}
}
