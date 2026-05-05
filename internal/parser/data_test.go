package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_matchFile(t *testing.T) {
	cwd, _ := os.Getwd()
	tests := []struct {
		name    string
		root    string
		pattern string
		want    int
	}{
		// TODO: Add test cases.
		{name: "sample", root: filepath.Join(cwd, "../../sample"),
			pattern: "*.yml", want: 1},
		{name: "testdata", root: filepath.Join(cwd, "../../testdata"),
			pattern: "*.yml", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := matchFile(tt.root, tt.pattern)
			if len(got) != tt.want {
				t.Fatalf("want %d got %v", tt.want, got)
			}
			t.Logf("%v", got)
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
