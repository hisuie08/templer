package output

import 	"io"

type stdoutOutput struct {
	w io.Writer
}

func (o *stdoutOutput) WriteFile(_ string, data []byte) error {
	_, err := o.w.Write(data)
	return err
}

func (o *stdoutOutput) IsStd() bool {
	return true
}
