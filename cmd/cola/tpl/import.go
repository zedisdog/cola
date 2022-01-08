package tpl

import (
	"bytes"
	"text/template"
)

const ImportTemp = `import{{if lt (len .) 2}} "{{- (index . 0).Import -}}" {{else}} (
{{range .}}	{{if .Alias}}{{.Alias}} {{end}}"{{.Import}}"
{{end}}){{end}}`

type ImportOptions struct {
	Alias  string
	Import string
}

func GenImport(s []ImportOptions) (r []byte, err error) {
	tmp, err := template.New("test").Parse(ImportTemp)
	if err != nil {
		return
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, s)
	if err != nil {
		return
	}
	r = buff.Bytes()
	return
}
