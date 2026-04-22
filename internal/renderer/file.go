package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"templer/internal/option"
)

func RenderFile(opt option.Option, data map[string]any) error {
	outDir := func() string {
		if opt.OutArg == option.OutSibling {
			return filepath.Dir(opt.Template.Value)
		}
		return opt.OutArg
	}()

	b, err := os.ReadFile(opt.Template.Value)
	if err != nil {
		fmt.Errorf(
			"cannot read path: %s: %w\nhint: use --literal to treat it as a string",
			opt.Template.Value, err,
		)
	}
	tmplStr := string(b)
	outPath := filepath.Join(outDir, strings.TrimSuffix(
		filepath.Base(opt.Template.Value), opt.Template.Suffix))

	var w = os.Stdout
	if opt.OutArg != "" {
		if i,err:=os.Stat(outPath);err==nil{
			if i.IsDir(){
				outPath=filepath.Join(outPath,filepath.Base(opt.Template.Value))
			}
		}
		w, err = createFile(outPath)
		if err != nil {
			return err
		}
		defer w.Close()
	}
	if opt.OutArg == "" {
		tmplStr = fmt.Sprintf("%s\n%s", outPath, tmplStr)
		if !strings.HasSuffix(tmplStr, "\n") {
			tmplStr = fmt.Sprintln(tmplStr)
		}
	}
	return render("", tmplStr, data, w)
}
