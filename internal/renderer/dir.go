package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"templer/internal/context"
	"templer/internal/option"
)

func (d *dirRenderer) fixPath() (string, error) {
	if d.opt.OutArg != "" {
		dir, err := isValidDir(d.opt.OutArg)
		if err != nil {
			return "", err
		}
		if !dir {
			return "", fmt.Errorf(
				"invalid path: %s is not a directory", d.opt.OutArg)
		}
		return filepath.Abs(d.opt.OutArg)
	}
	if d.opt.OutDefault {
		return d.opt.Template.Value, nil
	}
	d.ctx.Out.AsStd()
	return "", nil
}

type dirRenderer struct {
	ctx  context.Context
	opt  option.Option
	data map[string]any
}

func (d *dirRenderer) Render() error {
	outDir, err := d.fixPath()
	if err != nil {
		return err
	}
	return filepath.WalkDir(d.opt.Template.Value,
		func(path string, e os.DirEntry, err error) error {
			return d.execEntry(path, e, err, outDir)
		})
}

func (d *dirRenderer) execEntry(
	path string, e os.DirEntry, err error, outDir string) error {
	if err != nil {
		return err
	}
	if e.IsDir() {
		return nil
	}
	if !strings.HasSuffix(path, d.opt.Template.Suffix) {
		return nil
	}
	rel, err := filepath.Rel(d.opt.Template.Value, path)
	if err != nil {
		return err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	tmplStr := string(b)
	outPath := filepath.Join(outDir, strings.TrimSuffix(
		rel, d.opt.Template.Suffix))
	return render(path, tmplStr, outPath, d.data, d.ctx.Out)
}
