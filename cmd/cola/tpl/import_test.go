package tpl

import (
	"bytes"
	"testing"
	"text/template"
)

func TestImportTemp(t *testing.T) {
	tmp, err := template.New("test").Parse(ImportTemp)
	if err != nil {
		t.Fatal(err)
	}
	buff := bytes.NewBuffer([]byte{})
	err = tmp.Execute(buff, []ImportOptions{
		{
			Alias:  "_",
			Import: "test",
		},
		{
			Import: "test1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	println(buff.String())
}
