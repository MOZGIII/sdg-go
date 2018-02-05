package sdg

import (
	"os"
	"text/template"
	"time"
)

// MakeTemplate constrcuts a template.
func MakeTemplate(tpl string) (*template.Template, error) {
	return template.New("file").Parse(tpl)
}

// Debug switches the global debug mode.
// In debug mode generated files will not be removed on error.
var Debug = false

// WriteTempalteOpenFlags are flags to be used for OpenFile at WriteTempalte.
var WriteTempalteOpenFlags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC

// WriteTempalteArgs holds argments for WriteTempalte func.
type WriteTempalteArgs struct {
	Rows     [][]string
	DstPath  string
	Template *template.Template // optional, defaults to a template made from DefaultTemplateString
	Params   interface{}
	Package  string // optional, defaults to DefaultPackage
}

// WriteTempalte writes the template to a file.
func WriteTempalte(args WriteTempalteArgs) error {
	if err := applyDefaults(&args); err != nil {
		return err
	}

	f, err := os.OpenFile(args.DstPath, WriteTempalteOpenFlags, 0600)
	if err != nil {
		return err
	}
	defer f.Close() // nolint: errcheck

	err = args.Template.Execute(f, struct {
		Timestamp time.Time
		Package   string
		Rows      [][]string
		Params    interface{}
	}{
		Timestamp: time.Now(),
		Package:   args.Package,
		Rows:      args.Rows,
		Params:    args.Params,
	})

	if err != nil {
		if !Debug {
			os.Remove(args.DstPath) // nolint: errcheck,gas
		}
		return err
	}
	return nil
}

func applyDefaults(args *WriteTempalteArgs) error {
	if args == nil {
		return nil
	}

	if args.Package == "" {
		args.Package = DefaultPackage
	}

	if args.Template == nil {
		tpl, err := MakeTemplate(DefaultTemplateString)
		if err != nil {
			return err
		}
		args.Template = tpl
	}

	return nil
}
