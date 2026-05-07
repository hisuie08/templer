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

func (p *parser) loadEnv() error {
	envs := os.Environ()
	for _, s := range envs {
		if err := p.setKV(s); err != nil {
			if errors.Is(err, ErrInvalidFormat) {
				return fmt.Errorf("%w in env", err)
			}
			return err
		}
	}
	return nil
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
				matches, err := matchFile(arg)
				if err != nil {
					return err
				}
				if len(matches) == 0 {
					return fmt.Errorf("no matches for pattern: %s\n[hint] consider use 'str:' prefix", arg)
				}
				if err := p.asGlob(arg); err != nil {
					return err
				}
			} else {
				// 3. 最後に文字列
				if err := p.asStr(arg); err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func (p *parser) loadSets(sets []string) error {
	for _, s := range sets {
		if f, err := os.Open(s); err == nil {
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				if err := p.setKV(scanner.Text()); err != nil {
					if errors.Is(err, ErrInvalidFormat) {
						return fmt.Errorf("%w in %s", err, s)
					}
					return err
				}
			}
		} else {
			if err := p.setKV(s); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *parser) Parse() (map[string]any, error) {
	if !p.opt.IgnoreEnv {
		if err := p.loadEnv(); err != nil {
			return p.data, err
		}
	}
	if err := p.loadData(p.opt.DataArgs); err != nil {
		return p.data, err
	}
	if err := p.loadSets(p.opt.SetValues); err != nil {
		return p.data, err
	}
	return p.data, nil
}

func New(opt option.Option, ctx context.Context) *parser {
	return &parser{ctx: ctx, data: map[string]any{}, opt: opt}
}
