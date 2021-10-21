package pather

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// New 根据给定根路径创建Pather实例
func New(root string) *Pather {
	abs, err := filepath.Abs(strings.TrimRight(root, "/"))
	if err != nil {
		panic(err)
	}
	return &Pather{
		root: abs,
	}
}

// NewProjectPath 尝试创建Cola项目根目录的路劲
func NewProjectPath() *Pather {
	var err error
	path := ""
	num := 0
	for {
		path, err = filepath.Abs(path)
		if err != nil {
			break
		}
		if _, err = os.Stat(fmt.Sprintf("%s/go.mod", path)); err != nil {
			if num == 20 {
				break
			}
			path = fmt.Sprintf("%s/../", path)
		} else {
			return New(path)
		}
		num += 1
	}

	return nil
}

type Pather struct {
	root string
}

//Gen generate path with given path
func (p Pather) Gen(path string) string {
	return fmt.Sprintf("%s/%s", p.root, strings.TrimLeft(path, "/"))
}

//Dir return dir name of path
func (p Pather) Dir(path string) string {
	return filepath.Dir(p.Gen(path))
}

func (p Pather) Instance(path string) *Pather {
	return New(p.Gen(path))
}
