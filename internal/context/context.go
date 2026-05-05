package context

import (
	"io"
	"templer/internal/output"
)

type Context struct {
	Root string
	Out output.Output
	Log io.Writer
	Err io.Writer
}
