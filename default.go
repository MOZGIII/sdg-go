package sdg

// DefaultTemplateString is the template string used if no other
// template string is specified.
const DefaultTemplateString = `// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package generated

{{ .Params.Preface }}

var {{ .Params.Var }} = {{ .Params.Type }}{
{{- range .Rows }}
    {{ call $.Params.ValFn . }},
{{- end }}
}
`
