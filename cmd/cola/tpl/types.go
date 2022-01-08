package tpl

import (
	"bytes"
	"github.com/zedisdog/cola/cmd/cola/tpl/funcs"
	"strings"
	"text/template"
)

const StructTemp = `type {{.Name}} struct { {{- range .Fields}}
	{{.Name}}{{if .Type}}	{{.Type}}{{end}}` + "{{if .Tags}}	`" + `{{range .Tags}}{{.Key}}:"{{.Value}}" {{end}}` + "`{{end}}" + `{{end}}
}
{{range .Methods}}
` + FuncTemp + `
{{end}}`

type Tag struct {
	Key   string
	Value string
}

type Field struct {
	Name string
	Type string
	Tags []Tag
}

type StructTempOptions struct {
	Name    string
	Fields  []Field
	Methods []FuncTempOptions
}

func GenStruct(options StructTempOptions) (r []byte, err error) {

	for i := range options.Methods {
		if options.Methods[i].Receiver == "" {
			options.Methods[i].Receiver = strings.ToLower(string(options.Name[0]))
		}
		if options.Methods[i].ReceiverType == "" {
			options.Methods[i].ReceiverType = options.Name
		}
	}

	tmp, err := template.New("struct").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
	}).Parse(StructTemp)
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

const InterfaceTemp = `type {{.Name}} interface { {{- range .Funcs}}
	{{.Name}}({{pList .Params}}) {{rList .Returns}}{{end}}
}`

type InterfaceMethod struct {
	Name   string
	Params map[string]string
	//Returns map[string]string or []string
	Returns interface{}
}

type InterfaceTempOptions struct {
	Name  string
	Funcs []InterfaceMethod
}

func GenInterface(options InterfaceTempOptions) (r []byte, err error) {
	tmp, err := template.New("interface").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
		"rList": funcs.ReturnList,
	}).Parse(InterfaceTemp)
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
