package swagger

import (
	"bytes"
	"embed"
	"io/fs"
	"io/ioutil"
	"text/template"
)

//go:embed dist/*
var dist embed.FS

type swaggerFs struct {
	fs   fs.FS
	spec string
}

func (s swaggerFs) Open(name string) (f fs.File, err error) {
	if name == "index.html" {
		var index fs.File
		index, err = s.fs.Open(name)
		if err != nil {
			return
		}
		var content []byte
		var tmpl *template.Template
		var result bytes.Buffer
		content, err = ioutil.ReadAll(index)
		tmpl, err = template.New("tpl").Parse(string(content))
		if err != nil {
			return
		}
		err = tmpl.Execute(&result, s.spec)
		if err != nil {
			return
		}
		f = &fakeFile{
			name:     "index.html",
			contents: result.String(),
			mode:     fs.ModePerm,
		}
		return
	}
	return s.fs.Open(name)
}

//SwaggerUI return an instance of fs.FS which
func SwaggerUI(specUrl string) fs.FS {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	return swaggerFs{
		fs:   sub,
		spec: specUrl,
	}
}
