package engine

import "os"

func loadTemplate(arg string) string {

	if b, err := os.ReadFile(arg); err == nil {
		return string(b)
	}

	return arg
}