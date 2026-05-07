package funcs

import (
	"os"
	"os/exec"
	"path/filepath"
)

func readFile(path string) (string, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func cwd() (string, error) {
	return os.Getwd()
}

func shell(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}
