package funcs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
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

const ShellExecWarning = `
[WARNING] The 'Exec' function executes external commands from within templates
  For safety:
    1. Allow only the minimum necessary commands
    2. Ensure you understand the contents of the <template>
    3. Run in your own risk`

type limitedBuffer struct {
	buf      bytes.Buffer
	maxBytes int64
	written  int64
	exceeded bool
}

func (b *limitedBuffer) Len() int {
	return b.buf.Len()
}

func (b *limitedBuffer) Write(p []byte) (int, error) {
	remain := b.maxBytes - b.written

	if remain <= 0 {
		b.exceeded = true
		return 0, ErrOutputLimitExceeded
	}

	if int64(len(p)) > remain {
		p = p[:remain]
		b.exceeded = true
	}

	n, err := b.buf.Write(p)
	b.written += int64(n)

	if b.exceeded {
		return n, ErrOutputLimitExceeded
	}

	return n, err
}

func (b *limitedBuffer) String() string {
	return b.buf.String()
}

func (t *TemplerFunc) execShell(cmd string, args ...string) (string, error) {
	if !t.opt.AllowShellExecution {
		return "", ErrShellDisabled
	}
	if !slices.Contains(t.opt.AllowedShell, cmd) {
		return "", &ShellDisallowedError{Command: cmd}
	}

	ctx, cancel := context.WithTimeout(context.Background(), t.shellTimeout)
	defer cancel()

	var stdout limitedBuffer
	var stderr limitedBuffer

	stdout.maxBytes = 1024 * 1024
	stderr.maxBytes = 1024 * 1024

	c := exec.CommandContext(ctx, cmd, args...)
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return stdout.String(),
			&ShellTimeOutError{
				Command: cmd,
			}
	}
	if stdout.exceeded || stderr.exceeded {
		return stdout.String(),
			ErrOutputLimitExceeded
	}
	if err != nil {
		if stderr.Len() > 0 {
			return stdout.String(), fmt.Errorf("%w: %s", err, stderr.String())
		}
		return stdout.String(), err
	}

	return stdout.String(), nil
}
