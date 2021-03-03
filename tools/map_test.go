package tools

import "testing"

func TestStruct2Map(t *testing.T) {
	type test struct {
		Test string `json:"test"`
	}
	te := test{
		Test: "123",
	}
	r := Struct2Map(te)
	if r["test"] != "123" {
		t.Fatal("转换失败")
	}
	r = Struct2Map(&te)
	if r["test"] != "123" {
		t.Fatal("转换失败2")
	}
}
