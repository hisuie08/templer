package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"templer/internal/context"
	"templer/internal/option"
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

func (p *parser) loadEnv() {
	envs := os.Environ()
	for _, s := range envs {
		p.setKV(s)
	}
}

func (p *parser) loadData(args []string) error {
	for _, arg := range args {
		// 文字列プレフィックスを確認
		switch {
		case strings.HasPrefix(arg, option.Prefix.Str):
			v := strings.TrimPrefix(arg, option.Prefix.Str)
			if err := p.asStr(v); err != nil {
				return err
			}
		case strings.HasPrefix(arg, option.Prefix.Glob):
			v := strings.TrimPrefix(arg, option.Prefix.Glob)
			if err := p.asGlob(v); err != nil {
				return err
			}
		case strings.HasPrefix(arg, option.Prefix.File):
			v := strings.TrimPrefix(arg, option.Prefix.File)
			if err := p.asFile(v); err != nil {
				return err
			}
		default:
			// ファイルとして試す
			if err := p.asFile(arg); err == nil {
				continue
				// 「ファイルが無い」のみ許容（次の評価へ進行）
			} else {
				if !errors.Is(err, os.ErrNotExist) {
					return err
				}
			}
			if hasMeta(arg) {
				matches, err := matchFile(p.ctx.Root, arg)
				if err != nil {
					return err
				}
				if len(matches) == 0 {
					return fmt.Errorf("no matches for pattern: %s", arg)
				}
				if err := p.asGlob(arg); err != nil {
					return err
				}
				continue
			}

			// 3. 最後に文字列
			if err := p.asStr(arg); err != nil {
				return err
			}
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
	if err := p.loadData(p.opt.DataArgs); err != nil {
		return p.data, err
	}
	p.loadSets(p.opt.SetValues)
	return p.data, nil
}

func New(opt option.Option, ctx context.Context) *parser {
	return &parser{ctx: ctx, data: map[string]any{}, opt: opt}
}
