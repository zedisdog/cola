package tpl

import (
	"bytes"
	"github.com/zedisdog/cola/cmd/cola/tpl/funcs"
	"testing"
	"text/template"
)

func TestMethodTemp(t *testing.T) {
	tmp, err := template.New("test").Funcs(template.FuncMap{
		"pList": funcs.ParamList,
	}).Parse(FuncTemp)
	if err != nil {
		t.Fatal(err)
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, MethodTempParams{
		Receiver:     "",
		ReceiverType: "test",
		Name:         "",
		Params: map[string]string{
			"a": "int",
			"b": "string",
		},
		Returns: []string{
			"string",
			"error",
		},
		Content: ``,
	})
	if err != nil {
		t.Fatal(err)
	}
	println(buff.String())
}
