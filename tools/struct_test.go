package tools

import (
	"testing"
)

func TestCopyFields(t *testing.T) {
	type a struct {
		Test2 string
	}
	type b struct {
		Test2 string
	}
	var (
		aa a
		bb b
	)
	aa.Test2 = "111"
	if err := CopyFields(aa, &bb); err != nil {
		t.Fatal(err.Error())
	} else {
		if bb.Test2 != "111" {
			t.Fatal("copy failed")
		}
	}
	bb.Test2 = ""
	if err := CopyFields(&aa, &bb); err != nil {
		t.Fatal(err.Error())
	} else {
		if bb.Test2 != "111" {
			t.Fatal("copy failed")
		}
	}
}
