package tpl

import "bytes"

func GenLineBreak(num int) (r []byte) {
	return bytes.Repeat([]byte("\n"), num)
}
