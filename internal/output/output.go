package output

import "io"

type Output interface {
	IsStd() bool
	WriteFile(path string, data []byte) error
}

func StdOut(w io.Writer) Output {
	return &stdoutOutput{w: w}
}

func FileOut() Output {
	return &fileOutput{}
}
