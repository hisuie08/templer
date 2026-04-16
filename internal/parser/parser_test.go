package parser

import (
	"reflect"
	"testing"
)

func Test_mapSetter(t *testing.T) {
	tests := []struct {
		name  string
		data  map[string]any
		key   string
		value any
		want  map[string]any
	}{
		{name: "simple str", data: map[string]any{}, key: "key1", value: "value1", want: map[string]any{"key1": "value1"}},
		{name: "simple int", data: map[string]any{}, key: "key1", value: 123, want: map[string]any{"key1": 123}},
		{name: "simple override", data: map[string]any{"key1": "value1"}, key: "key1", value: "value2", want: map[string]any{"key1": "value2"}},
		{name: "nested str", data: map[string]any{}, key: "nested.key1", value: "value1", want: map[string]any{"nested": map[string]any{"key1": "value1"}}},
		{name: "nested override", data: map[string]any{"nested": "exist"}, key: "nested.key1", value: "value1", want: map[string]any{"nested": map[string]any{"key1": "value1"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapSetter(tt.data, tt.key, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapSetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
