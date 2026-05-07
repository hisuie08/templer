package output

import (
	"io"
	"os"
	"path/filepath"
)

type Output interface {
	IsStd() bool
	WriteFile(path string, data []byte) error
	AsStd()
}

func createFile(path string) (*os.File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
		return nil, err
	}
	return os.Create(abs)
}

type outCtler struct {
	isStdOut     bool
	stdOutWriter io.Writer
}

func (o *outCtler) IsStd() bool {
	return o.isStdOut
}
func (o *outCtler) WriteFile(path string, data []byte) error {
	if o.IsStd() {
		_, err := o.stdOutWriter.Write(data)
		return err
	}
	p, err := createFile(path)
	if err != nil {
		return err
	}
	defer p.Close()
	if _, err := p.Write(data); err != nil {
		return err
	}
	return nil
}

func (o *outCtler) AsStd() {
	o.isStdOut = true
}

func OutController(w io.Writer) Output {
	return &outCtler{stdOutWriter: w, isStdOut: false}
}
