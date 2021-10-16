package database

import (
	"testing"
)

func TestGetDialector(t *testing.T) {
	_, err := getDialector("mysql://root:toor@tcp(localhost)/main")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWithDsn(t *testing.T) {
	o := &Options{}
	WithDsn("mysql://root:toor@tcp(localhost)/main?abc=a/b&abc=c/d")(o)
	if o.dsn != "mysql://root:toor@tcp(localhost)/main?abc=a%2Fb&abc=c%2Fd" {
		println(o.dsn)
		t.Fatal("fail")
	}
}
