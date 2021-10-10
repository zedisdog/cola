package lib

import (
	"bytes"
	"fmt"
	"github.com/zedisdog/cola/cmd/cola/stubs"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Export(path string, moduleName string) (err error) {
	dest := strings.TrimRight(path, "\\/")

	info, err := os.Stat(dest)
	if err == nil {
		var (
			file  *os.File
			names []string
		)

		if info.IsDir() {
			file, err = os.Open(dest)
			if err != nil {
				return
			}
			names, err = file.Readdirnames(-1)
			if err != nil {
				return
			}
			if len(names) > 0 {
				fmt.Printf("the folder \"%s\" is not empty, files might be overwrote, continue? [y/n]: ", dest)
				var yesOrNo string
				fmt.Scan(&yesOrNo)
				if !strings.EqualFold("y", strings.TrimSpace(yesOrNo)) {
					return
				}
			}
			file.Close()
		} else {
			fmt.Printf("the file \"%s\" is exists, file might be overwrote, continue? [y/n]: ", dest)
			var yesOrNo string
			fmt.Scan(&yesOrNo)
			if !strings.EqualFold("y", strings.TrimSpace(yesOrNo)) {
				return
			}
		}
	} else {
		err = nil
	}
	os.RemoveAll(dest)
	os.MkdirAll(dest, 0755)

	fs.WalkDir(stubs.Template, "template", func(path string, d fs.DirEntry, err error) (e error) {
		if d.IsDir() || err != nil {
			return
		}
		destPath, content, err := convertTemplateToSource(dest, "template", path, moduleName, stubs.Template)
		if err != nil {
			return
		}
		err = os.MkdirAll(filepath.Dir(destPath), 0755)
		if err != nil {
			return
		}
		err = os.WriteFile(destPath, content, 0744)
		return
	})
	return
}

func convertTemplateToSource(destDir string, srcDir string, srcPath string, moduleName string, f fs.FS) (destPath string, content []byte, err error) {
	destPath = makeExportPath(destDir, srcDir, srcPath)
	content, err = fs.ReadFile(f, srcPath)
	if err != nil {
		return
	}
	if filepath.Ext(srcPath) == ".stub" {
		var tmpl *template.Template
		tmpl, err = template.New("t").Parse(string(content))
		if err != nil {
			return
		}
		buffer := bytes.NewBuffer([]byte{})
		err = tmpl.Execute(buffer, moduleName)
		content = buffer.Bytes()
	}
	return
}

func makeExportPath(destDir string, srcDir string, srcPath string) (path string) {
	path = strings.Replace(srcPath, srcDir, destDir, 1)
	if filepath.Ext(path) == ".stub" { // convert stub file to go source code file
		path = strings.Replace(path, ".stub", "", 1)
	}
	return
}
