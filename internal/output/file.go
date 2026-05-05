package output

import (
	"os"
	"path/filepath"
)

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

type fileOutput struct{}

func (o *fileOutput) IsStd() bool {
	return false
}
func (o *fileOutput) WriteFile(path string, data []byte) error {
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
