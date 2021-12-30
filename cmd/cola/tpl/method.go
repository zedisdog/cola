package tpl

import (
	"bytes"
	"github.com/zedisdog/cola/cmd/cola/tpl/funcs"
	"text/template"
)

type MethodTempParams struct {
	Receiver     string
	ReceiverType string
	Name         string
	Params       map[string]string
	//Returns map[string]string or []string
	Returns interface{}
	Content string
}

const FuncTemp = `func{{if .Name}} {{if .Receiver}}({{.Receiver}} {{.ReceiverType}}) {{end}}{{.Name}}{{end}}({{pList .Params}}) ({{pList .Returns}}) { {{- if .Content}}
	{{.Content}}
{{end -}} }
`

func GenMethod(params MethodTempParams) (r []byte, err error) {
	tmp, err := template.New("method").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
	}).Parse(FuncTemp)
	if err != nil {
		return
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, params)
	if err != nil {
		return
	}
	r = buff.Bytes()
	return
}
