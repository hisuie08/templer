package funcs

import (
	"os"
	"path/filepath"
)

func readFile(path string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	p := filepath.Join(wd, path)
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func cwd() (string, error) {
	return os.Getwd()
}
