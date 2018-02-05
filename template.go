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
	Template *template.Template
	Params   interface{}
}

// WriteTempalte writes the template to a file.
func WriteTempalte(args WriteTempalteArgs) error {
	f, err := os.OpenFile(args.DstPath, WriteTempalteOpenFlags, 0600)
	if err != nil {
		return err
	}
	defer f.Close() // nolint: errcheck

	err = args.Template.Execute(f, struct {
		Timestamp time.Time
		Rows      [][]string
		Params    interface{}
	}{
		Timestamp: time.Now(),
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
