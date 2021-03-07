package drivers

import (
	"errors"
	"github.com/zedisdog/cola/pather"
	"io"
	"os"
)

func NewLocal(path string) *Local {
	p := pather.New(path)
	err := os.MkdirAll(p.Gen(""), 0766)
	if err != nil {
		panic(err)
	}
	return &Local{
		path: pather.New(path),
	}
}

type Local struct {
	path *pather.Pather
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
