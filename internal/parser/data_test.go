package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
)

var fsMap = map[string][]string{
	"match1":    []string{"file1.txt", "file2.yml"},
	"match0":    []string{"file1.txt", "file2.txt"},
	"match.yml": []string{"file1.txt", "file2.txt"},
	"matchs":    []string{"file.yml", "test.yml"},
}

func mockFs(root string) error {
	for dirname, v := range fsMap {
		dir := filepath.Join(root, dirname)
		if e := os.MkdirAll(dir, 0755); e != nil {
			return e
		}
		for _, f := range v {
			file := filepath.Join(dir, f)
			if e := os.WriteFile(file, []byte("0"), 0644); e != nil {
				return e
			}
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
		{name: "nest", dir: ".", want: 3},
	}
	if err := mockFs(td); err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := filepath.Join(td, tt.dir)
			got, err := matchFile(dir, "**/*.yml")
			if err != nil {
				t.Fatal(err)
			}
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

func Test_sortPath(t *testing.T) {
	tests := []struct {
		name  string
		paths []string
		want  []string
	}{{
		name: "test",
		paths: []string{
			"b/c/d.txt", "a.txt", "b/a.txt", "b/c.txt", "a/b/c.txt", "a/b.txt"},
		want: []string{
			"a.txt", "a/b.txt", "b/a.txt", "b/c.txt", "a/b/c.txt", "b/c/d.txt"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortPath(tt.paths)
			eq := deep.Equal(got, tt.want)
			if len(eq) != 0 {
				t.Errorf("%v", eq)
			}
		})
	}
}
