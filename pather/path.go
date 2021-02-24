package pather

import (
	"fmt"
	"strings"
)

func New(root string) *Pather {
	return &Pather{
		root: strings.TrimRight(root, "/"),
	}
}

type Pather struct {
	root string
}

func (p Pather) Gen(path string) string {
	return fmt.Sprintf("%s/%s", p.root, strings.TrimLeft(path, "/"))
}
