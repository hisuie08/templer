package parser

import (
	"errors"
	"fmt"
	"strings"
)

// setKvはKey=Value型のデータを処理する
// loadEnvとloadSetで使う
var ErrInvalidFormat = errors.New("invalid format")

type errInvalidSet struct {
	Set string
}

func (e *errInvalidSet) Error() string {
	return fmt.Sprintf("invalid format: %s", e.Set)
}
func (e *errInvalidSet) Unwrap() error {
	return ErrInvalidFormat
}

func (p *parser) setKV(s string) error {
	kv := strings.SplitN(s, "=", 2)
	if len(kv) == 2 {
		mapSetter(p.data, kv[0], kv[1])
		return nil
	}
	return &errInvalidSet{Set: s}
}

// テスト用の切り出し
func mapSetter(data map[string]any, key string, value any) map[string]any {
	selector := strings.Split(key, ".")
	current := data
	for i, key := range selector {
		if i == len(selector)-1 {
			current[key] = value
			return data
		}
		if next, ok := current[key]; ok {
			if m, ok := next.(map[string]any); ok {
				current = m
				continue
			}
		}
		newMap := map[string]any{}
		current[key] = newMap
		current = newMap
	}
	return data
}
