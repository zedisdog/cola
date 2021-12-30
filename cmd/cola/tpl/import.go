package tpl

import (
	"bytes"
	"text/template"
)

const ImportTemp = `import{{if lt (len .) 2}} "{{index . 0}}" {{else}} (
{{range .}}	"{{.}}"
{{end}}){{end}}`

func GenImport(s []string) (r []byte, err error) {
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
