package renderer

import (
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	td := t.TempDir()
	testCases := []struct {
		desc string
		path string
	}{
		{desc: "current", path: "./test.txt"},
		{desc: "parent", path: "../test.txt"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a, _ := filepath.Abs(filepath.Join(td, tC.path))
			dir := filepath.Dir(a)
			file := filepath.Base(a)
			t.Logf("full: %s", a)
			t.Logf("dir: %s", dir)
			t.Logf("file: %s", file)

		})
	}
}
