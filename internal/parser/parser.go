package parser

import (
	"bufio"
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
		p.setKV(s)
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
		if f, err := os.Open(s); err == nil {
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				p.setKV(scanner.Text())
			}

		} else {
			p.setKV(s)
		}
	}
}

func (p *parser) setKV(s string) {
	kv := strings.SplitN(s, "=", 2)
	if len(kv) == 2 {
		mapSetter(p.data, kv[0], kv[1])
	}
}

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
