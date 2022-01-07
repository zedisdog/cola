package tpl

import (
	"bytes"
	"github.com/zedisdog/cola/cmd/cola/tpl/funcs"
	"strings"
	"text/template"
)

const FileTemp = `package {{.PkgName}}{{if .Imports}}

{{template "imports" .Imports}}{{end}}{{- range .Interfaces}}

` + InterfaceTemp + `
{{end}}{{range .Funcs}}
` + FuncTemp + `{{end}}{{range .Structs}}

` + StructTemp + `{{end}}{{- define "imports"}}` + ImportTemp + `{{end -}}`

type FileTempOptions struct {
	PkgName    string
	Imports    []ImportOptions
	Interfaces []InterfaceTempOptions
	Funcs      []FuncTempOptions
	Structs    []StructTempOptions
}

func GenFile(options FileTempOptions) (r []byte, err error) {
	for i := range options.Structs {
		for j := range options.Structs[i].Methods {
			if options.Structs[i].Methods[j].Receiver == "" {
				options.Structs[i].Methods[j].Receiver = strings.ToLower(string(options.Structs[i].Name[0]))
			}
			if options.Structs[i].Methods[j].ReceiverType == "" {
				options.Structs[i].Methods[j].ReceiverType = options.Structs[i].Name
			}
		}
	}

	tmp, err := template.New("file").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
		"rList": funcs.ReturnList,
	}).Parse(FileTemp)
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
