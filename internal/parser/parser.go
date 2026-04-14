package parser

import (
	"encoding/json"
	"os"
	"strings"
	"templer/internal/option"

	"gopkg.in/yaml.v3"
)

type parser struct {
	opt  option.Option
	data map[string]any
}

func (p *parser) loadEnv() {
	envs := os.Environ()
	for _, s := range envs {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) == 2 {
			p.data[kv[0]] = kv[1]
		}
	}
}

func (p *parser) loadData(args []string, format string) {
	raw := ""
	for _, arg := range args {
		// Read the contents if it's a readable file
		if b, err := os.ReadFile(arg); err == nil {
			raw += "\n" + string(b)
		} else {
			// Continue as stdin string
			raw += "\n" + arg
		}
		if raw != "" {
			if format == "json" {
				json.Unmarshal([]byte(raw), &p.data)
			} else {
				yaml.Unmarshal([]byte(raw), &p.data)
			}
		}
	}
}

func (p *parser) loadSets(sets []string) {
	for _, s := range sets {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) == 2 {
			p.data[kv[0]] = kv[1]
		}
	}
}

func (p *parser) Parse() map[string]any {
	if p.opt.LoadEnv {
		p.loadEnv()
	}
	p.loadData(p.opt.DataArgs, p.opt.DataFormat)
	p.loadSets(p.opt.SetValues)
	return p.data
}

func New(opt option.Option) *parser {
	return &parser{data: map[string]any{}, opt: opt}
}
func Load(args []string, format string, sets []string, loadEnv bool) (
	map[string]any, error) {

	data := map[string]any{}
	// read environments
	if loadEnv {
		envs := os.Environ()
		for _, s := range envs {
			kv := strings.SplitN(s, "=", 2)
			if len(kv) == 2 {
				data[kv[0]] = kv[1]
			}
		}
	}
	// load data
	raw := ""
	for _, arg := range args {
		// Read the contents if it's a readable file
		if b, err := os.ReadFile(arg); err == nil {
			raw += "\n" + string(b)
		} else {
			// Continue as stdin string
			raw += "\n" + arg
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
