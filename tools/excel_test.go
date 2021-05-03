package tools

import (
	"testing"
)

func TestCol(t *testing.T) {
	s := ColIndexByNum(300)

	if s != "KN" {
		t.Fatal("error")
	}
}
