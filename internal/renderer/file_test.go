package renderer

import (
	"os"
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

func TestStat(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("not join")
	t.Log(filepath.Clean("../dir.go"))
	t.Log("join")
	t.Log(filepath.Join(wd, "/dir.go"))
}
