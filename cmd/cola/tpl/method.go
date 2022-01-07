package tpl

import (
	"bytes"
	"github.com/zedisdog/cola/cmd/cola/tpl/funcs"
	"text/template"
)

type FuncTempOptions struct {
	Receiver     string
	ReceiverType string
	Name         string
	Params       map[string]string
	//Returns map[string]string or []string
	Returns interface{}
	Content string
}

const FuncTemp = `func{{if .Name}} {{if .Receiver}}({{.Receiver}} {{.ReceiverType}}) {{end}}{{.Name}}{{end}}({{pList .Params}}) {{rList .Returns}} { {{- if .Content}}
	{{.Content}}
{{end -}} }`

func GenFunc(options FuncTempOptions) (r []byte, err error) {
	tmp, err := template.New("func").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
		"rList": funcs.ReturnList,
	}).Parse(FuncTemp)
	if err != nil {
		return
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, options)
	if err != nil {
		return
	}
	r = buff.Bytes()
	return
}

type FuncSignTempOptions struct {
	Receiver     string
	ReceiverType string
	Name         string
	Params       map[string]string
	//Returns map[string]string or []string
	Returns interface{}
}

const FuncSignTemp = `func{{if .Name}} {{if .Receiver}}({{.Receiver}} {{.ReceiverType}}) {{end}}{{.Name}}{{end}}({{pList .Params}}) {{rList .Returns}}
`

func GenFuncSign(options FuncSignTempOptions) (r []byte, err error) {
	tmp, err := template.New("funcSign").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
		"rList": funcs.ReturnList,
	}).Parse(FuncSignTemp)
	if err != nil {
		return
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, options)
	if err != nil {
		return
	}
	r = buff.Bytes()
	return
}
