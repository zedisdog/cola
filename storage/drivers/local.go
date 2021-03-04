package drivers

import "github.com/zedisdog/cola/pather"

type Local struct {
	path *pather.Pather
}

func (l Local) Put(path string, data []byte) error {

	panic("implement me")
}

func (l Local) Get(path string) ([]byte, error) {
	panic("implement me")
}

func (l Local) Remove(path string) error {
	panic("implement me")
}
