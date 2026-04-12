package data

import (
	"encoding/json"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func Load(args []string, format string, sets []string) (map[string]any, error) {

	data := map[string]any{}
	envs := os.Environ()
	for _, s := range envs {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) == 2 {
			data[kv[0]] = kv[1]
		}
	}
	raw := ""
	for _, arg := range args {
		if b, err := os.ReadFile(arg); err == nil {
			raw += "\n" + string(b)
		}
		if raw != "" {
			if format == "json" {
				json.Unmarshal([]byte(raw), &data)
			} else {

				yaml.Unmarshal([]byte(raw), &data)
			}
		}
	}

	for _, s := range sets {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) == 2 {
			data[kv[0]] = kv[1]
		}
	}

	return data, nil
}
