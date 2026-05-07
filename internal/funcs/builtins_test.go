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
			got, gotErr := readFile(tt.path)
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
