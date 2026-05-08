package funcs

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func (t *TemplerFunc) readFile(path string) (string, error) {
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

func (t *TemplerFunc) cwd() (string, error) {
	return os.Getwd()
}

// TODO: 実行可能コマンドのセキュリティ
func (t *TemplerFunc) execShell(cmd string, args ...string) (string, error) {
	if !t.opt.AllowShellExecution {
		return "", errors.New(
			"function Exec not permitted\nuse --allow-shell-execution")
	}
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}
