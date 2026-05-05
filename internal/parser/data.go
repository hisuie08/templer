package parser

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"templer/internal/option"
)

// --data [arg] 処理用関数群
type GlobNoMatchError struct {
	Pattern string
}

func (e *GlobNoMatchError) Error() string {
	return fmt.Sprintf("no matches for pattern: %s", e.Pattern)
}

func glob(path, pattern string) []string {
	g, err := filepath.Glob(filepath.Join(path, pattern))
	if err != nil {
		return []string{}
	}
	return g
}
func matchFile(root string, pattern string) []string {
	result := []string{}
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			result = append(result, glob(path, pattern)...)
		}
		return nil
	})
	return result
}

func hasMeta(str string) bool {
	return strings.ContainsAny(str, option.MetaString)
}

