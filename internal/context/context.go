package context

import (
	"io"
	"templer/internal/output"
)

type Context struct {
	Out output.Output
	Log io.Writer
	Err io.Writer
}
