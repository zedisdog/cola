package generate

import (
	"github.com/zedisdog/cola/cmd/cola/tpl"
	"os"
)

func File(options tpl.FileTempOptions, path string) (err error) {
	content, err := tpl.GenFile(options)
	if err != nil {
		return
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0766)
	defer f.Close()

	_, err = f.Write(content)
	return
}
