package storage

import (
	"sync"
	"testing"
)

var inst *a

type a struct {
	a string
}

func getA() a {
	var once sync.Once
	once.Do(func() {
		inst = &a{a: "123"}
	})
	return *inst
}

func TestNormal(t *testing.T) {
	a1 := getA()
	a2 := getA()
	a1.a = "111"
	if a2.a != "123" {
		t.Fatal("error")
	}
}
