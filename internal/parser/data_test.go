package parser

import (
	"os"
	"path/filepath"
	"testing"
)

var fsMap = map[string][]string{
	"match1":    []string{"file1.txt", "file2.yml"},
	"match0":    []string{"file1.txt", "file2.txt"},
	"match.yml": []string{"file1.txt", "file2.txt"},
}

func mockFs(root, dirname string) error {
	dir := filepath.Join(root, dirname)
	if e := os.MkdirAll(dir, 0755); e != nil {
		return e
	}
	for _, v := range fsMap[dirname] {
		file := filepath.Join(dir, v)
		if e := os.WriteFile(file, []byte("0"), 0644); e != nil {
			return e
		}
	}
	return nil
}
func Test_matchFile(t *testing.T) {
	td := t.TempDir()
	tests := []struct {
		name string
		dir  string
		want int
	}{
		{name: "match 1", dir: "match1", want: 1},
		{name: "match 0", dir: "match0", want: 0},
		{name: "ignore dir", dir: "match.yml", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mockFs(td, tt.dir); err != nil {
				t.Fatal(err)
			}
			dir := filepath.Join(td, tt.dir)
			got := matchFile(dir, "*.yml")
			if len(got) != tt.want {
				t.Fatalf("want %d got %v", tt.want, got)
			}
		})
	}
}

func Test_hasMeta(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		want bool
	}{
		{name: "n/a", str: "test", want: false},
		{name: "*", str: "*test", want: true},
		{name: "?", str: "?test", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasMeta(tt.str)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("hasMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
