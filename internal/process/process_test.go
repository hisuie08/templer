package process

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_getPath(t *testing.T) {
	wd, _ := os.Getwd()
	tests := []struct {
		name    string
		p       string
		want    string
		wantErr bool
	}{
		{name: "lcl1", p: "./test.txt", want: filepath.Join(wd, "test.txt")},
		{name: "lcl2", p: "test.txt", want: filepath.Join(wd, "test.txt")},
		{name: "lcl3", p: "../test.txt", want: filepath.Join(wd, "../test.txt")},
		{name: "abs", p: "/path/to/test.txt", want: "/path/to/test.txt"},
		{name: "pwd", p: "./", want: wd},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, _ := getOutPath(tt.p)
			if got != tt.want {
				t.Fatalf("unexpected return: %s want %s", got, tt.want)
			}

		})
	}
}
