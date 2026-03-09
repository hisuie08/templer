package input

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func Load(arg, format string, sets []string) (map[string]any, error) {

	data := map[string]any{}

	raw := ""

	if arg == "" {
		b, _ := io.ReadAll(os.Stdin)
		raw = string(b)
	} else if b, err := os.ReadFile(arg); err == nil {
		raw = string(b)
	} else {
		raw = arg
	}

	if raw != "" {
		if format == "json" {
			json.Unmarshal([]byte(raw), &data)
		} else {
			yaml.Unmarshal([]byte(raw), &data)
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
