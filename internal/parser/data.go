package parser

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"templer/internal/option"

	"github.com/bmatcuk/doublestar/v4"
	"gopkg.in/yaml.v3"
)

// --data [arg] 処理用関数群

// glob関数を再帰的に使ってマッチしたパスを返す
func matchFile(root string, pattern string) ([]string, error) {
	return doublestar.Glob(os.DirFS(root), pattern, doublestar.WithFilesOnly())
}

func hasMeta(s string) bool {
	return strings.ContainsAny(s, option.MetaStr)
}

// ディレクトリネスト昇順、名前昇順でソート
// globのマッチを安定させる
func sortPath(paths []string) []string {
	depth := func(path string) int {
		clean := filepath.Clean(path)
		if clean == "." {
			return 0
		}
		return strings.Count(clean, string(filepath.Separator))
	}
	sort.Slice(paths, func(i, j int) bool {
		di := depth(paths[i])
		dj := depth(paths[j])
		if di != dj {
			return di < dj
		}
		return paths[i] < paths[j]
	})
	return paths
}

func (p *parser) asGlob(v string) error {
	matches, err := matchFile(p.ctx.Root, v)
	if err != nil {
		return err
	}
	for _, v := range sortPath(matches) {
		if err := p.asFile(v); err != nil {
			return err
		}
	}
	return nil
}
func (p *parser) asFile(s string) error {
	v := filepath.Join(p.ctx.Root, s)
	b, err := os.ReadFile(v)
	if err != nil {
		return err
	}
	return p.parseArg(string(b))
}

func (p *parser) asStr(s string) error {
	return p.parseArg(s)
}

func (p *parser) parseArg(raw string) error {
	if p.opt.DataFormat == "json" {
		if err := json.Unmarshal([]byte(raw), &p.data); err != nil {
			return err
		}
	} else {
		if err := yaml.Unmarshal([]byte(raw), &p.data); err != nil {
			return err
		}
	}
	return nil
}
