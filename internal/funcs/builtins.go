package funcs

import (
	"errors"
	"fmt"
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

var ErrShellDisabled error = errors.New(
	"function 'Exec' disabled by default\nuse '--allow-shell-execution' to enable")

var ErrShellExecution error = &errShelExecution{}

type errShelExecution struct {
	Command string
}

func (e errShelExecution) Error() string {
	return fmt.Sprintf("disallowed command: %s is not in whitelist", e.Command)
}
func (e errShelExecution) Unwrap() error {
	return ErrShellExecution
}

// TODO: 実行可能コマンドのセキュリティ
func (t *TemplerFunc) execShell(cmd string, args ...string) (string, error) {
	if !t.opt.AllowShellExecution {
		return "", ErrShellDisabled
	}
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}
