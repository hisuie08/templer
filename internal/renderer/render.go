package renderer

import (
	"bytes"
	"os"
	"templer/internal/context"
	"templer/internal/funcs"
	"templer/internal/option"
	"templer/internal/output"
	"text/template"
	"text/template/parse"
)

func render(name, input string, outPath string,
	data map[string]any, out output.Output, opt option.Option) error {
	t, err := newTmpl(name, opt).Parse(input)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	if err = t.Execute(b, data); err != nil {
		return err
	}
	return out.WriteFile(outPath, fixForOut(b.Bytes(), out.IsStd()))
}

func newTmpl(name string, o option.Option) *template.Template {
	return template.New(name).Funcs(funcs.New(o).Funcs())
}
func fixForOut(b []byte, isStd bool) []byte {
	nl := []byte("\n")
	if !bytes.HasSuffix(b, nl) {
		b = append(b, nl...)
	}
	return b
}
func isValidDir(path string) (bool, error) {
	if i, err := os.Stat(path); err == nil {
		return i.IsDir(), nil
	} else {
		return false, err
	}
}

func hasFunc(tree *parse.Tree, name string) bool {
	var found bool

	var walk func(parse.Node)
	walk = func(n parse.Node) {
		if n == nil || found {
			return
		}

		switch x := n.(type) {
		case *parse.ListNode:
			for _, node := range x.Nodes {
				walk(node)
			}

		case *parse.ActionNode:
			walk(x.Pipe)

		case *parse.PipeNode:
			for _, cmd := range x.Cmds {
				walk(cmd)
			}

		case *parse.CommandNode:
			if len(x.Args) > 0 {
				if id, ok := x.Args[0].(*parse.IdentifierNode); ok && id.Ident == name {
					found = true
					return
				}
			}

		case *parse.IfNode:
			walk(x.Pipe)
			walk(x.List)
			walk(x.ElseList)

		case *parse.RangeNode:
			walk(x.Pipe)
			walk(x.List)
			walk(x.ElseList)

		case *parse.WithNode:
			walk(x.Pipe)
			walk(x.List)
			walk(x.ElseList)

		case *parse.TemplateNode:
			// 必要なら named template 辿る
		}
	}

	walk(tree.Root)
	return found
}

type Renderer interface {
	Render() error
}

func Literal(ctx context.Context,
	opt option.Option, data map[string]any) Renderer {
	return &literalRenderer{ctx: ctx, opt: opt, data: data}
}

func File(ctx context.Context,
	opt option.Option, data map[string]any) Renderer {
	return &fileRenderer{ctx: ctx, opt: opt, data: data}
}
func Dir(ctx context.Context,
	opt option.Option, data map[string]any) Renderer {
	return &dirRenderer{ctx: ctx, opt: opt, data: data}
}
