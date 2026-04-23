package outwriter

import (
	"io"
	"os"
)

type noopWriterCloser struct {
	io.Writer
}

func (noopWriterCloser) Close() error {
	return nil
}

func outWriter(w io.WriteCloser) io.Writer {
	return os.Stdout
}
