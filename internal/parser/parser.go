package parser

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"templer/internal/context"
	"templer/internal/option"

	"gopkg.in/yaml.v3"
)

type parser struct {
	ctx  context.Context
	opt  option.Option
	data map[string]any
}

// setKvはKey=Value型のデータを処理する
// loadEnvとloadSetで使う
func (p *parser) setKV(s string) {
	kv := strings.SplitN(s, "=", 2)
	if len(kv) == 2 {
		mapSetter(p.data, kv[0], kv[1])
	}
}

func (p *parser) parseArg(arg string, format string) {
	var raw string
	if b, err := os.ReadFile(arg); err == nil {
		raw = "\n" + string(b)
	} else {
		// Continue as stdin string
		raw = "\n" + arg
	}
	if raw != "" {
		if format == "json" {
			json.Unmarshal([]byte(raw), &p.data)
		} else {
			yaml.Unmarshal([]byte(raw), &p.data)
		}
	}
}

func (p *parser) loadEnv() {
	envs := os.Environ()
	for _, s := range envs {
		p.setKV(s)
	}
}

func (p *parser) loadData(args []string, format string) error {
	for _, arg := range args {
		// メタ文字を含む場合Globチャレンジ
		if hasMeta(arg) {
			matches := matchFile(p.ctx.Root, arg)
			if len(matches) == 0 {
				return &GlobNoMatchError{arg}
			}
			for _, v := range matches {
				p.parseArg(v, format)
			}
		} else {
			p.parseArg(arg, format)
		}
	}
	return nil
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

func (p *parser) Parse() (map[string]any, error) {
	if p.opt.LoadEnv {
		p.loadEnv()
	}
	if err := p.loadData(p.opt.DataArgs, p.opt.DataFormat); err != nil {
		return p.data, err
	}
	p.loadSets(p.opt.SetValues)
	return p.data, nil
}

func New(opt option.Option, ctx context.Context) *parser {
	return &parser{ctx: ctx, data: map[string]any{}, opt: opt}
}
