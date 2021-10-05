package tools

import (
	"errors"
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

func TestCopyFields2(t *testing.T) {
	a := struct {
		Name       string
		Code       string
		AccountID  int
		TypeID     int
		IsOfficial int
	}{
		AccountID:  2,
		Name:       "张哲的店铺",
		Code:       "111",
		TypeID:     1,
		IsOfficial: 0,
	}

	b := struct {
		AccountID  int
		TypeID     int
		Code       string
		Name       string
		IsOfficial int
	}{
		AccountID:  2,
		Name:       "张哲的店铺",
		Code:       "111",
		TypeID:     1,
		IsOfficial: 1,
	}

	err := CopyFields(a, &b)

	if err != nil {
		t.Fatal(err)
	}
	if b.IsOfficial != 0 {
		t.Fatal(errors.New("isOfficial should be 0"))
	}
}
