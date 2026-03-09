package files

import (
	"os"
)

func Funcs() map[string]any {
	return map[string]any{
		"readFile": readFile,
	}
}

func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(b)
}
