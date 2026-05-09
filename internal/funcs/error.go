package funcs

import (
	"errors"
	"fmt"
)

var ErrShellDisabled error = errors.New(
	`function 'Exec' disabled by default
	use '--allow-shell-execution' to enable`)

var ErrShellDisallowed = errors.New("disallowed shell command")

type ShellDisallowedError struct {
	Command string
}

func (e *ShellDisallowedError) Error() string {
	return fmt.Sprintf(
		"disallowed command: '%[1]s'\nuse '--allow-command %[1]s' to allow it",
		e.Command)
}

func (e *ShellDisallowedError) Unwrap() error {
	return ErrShellDisallowed
}

var ErrOutputLimitExceeded = errors.New(
	"command output limit exceeded",
)

var ErrShellTimeout = errors.New("command execution timed out")

type ShellTimeOutError struct {
	Command string
}

func (e *ShellTimeOutError) Error() string {
	return fmt.Sprintf("command execution timed out: '%s'", e.Command)
}

func (e *ShellTimeOutError) Unwrap() error {
	return ErrShellTimeout
}
