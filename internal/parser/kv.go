package parser

import "strings"

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
