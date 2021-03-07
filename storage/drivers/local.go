package drivers

import (
	"errors"
	"fmt"
	"github.com/h2non/filetype"
	"github.com/zedisdog/cola/pather"
	"io"
	"os"
)

func NewLocal(path string, baseUrl string) *Local {
	p := pather.New(path)
	err := os.MkdirAll(p.Gen(""), 0766)
	if err != nil {
		panic(err)
	}
	return &Local{
		path:    pather.New(path),
		baseUrl: baseUrl,
	}
}

type Local struct {
	path    *pather.Pather
	baseUrl string
}

func (l Local) Put(path string, data []byte) (err error) {
	err = os.Mkdir(l.path.Dir(path), 0766)
	if err != nil {
		return
	}
	f, err := os.OpenFile(l.path.Gen(path), os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(data)
	return
}

func (l Local) Get(path string) (data []byte, err error) {
	f, err := os.Open(l.path.Gen(path))
	if err != nil {
		return
	}
	defer f.Close()
	data, err = io.ReadAll(f)
	return
}

func (l Local) Remove(path string) (err error) {
	_, err = os.Stat(l.path.Gen(path))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return
	}

	err = os.Remove(l.path.Gen(path))
	return
}

func (l Local) Path(path string) string {
	return l.path.Gen(path)
}

func (l Local) GetUrl(path string) string {
	return fmt.Sprintf("%s/%s", l.baseUrl, path)
}

func (l Local) GetMime(path string) string {
	fp, err := os.Open(l.path.Gen(path))
	if err != nil {
		return ""
	}
	defer fp.Close()
	b := make([]byte, 262)
	if _, err := fp.Read(b); err != nil {
		return ""
	}
	kind, _ := filetype.Match(b)
	if kind == filetype.Unknown {
		return ""
	}

	return kind.MIME.Value
}
