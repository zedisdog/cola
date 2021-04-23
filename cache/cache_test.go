package cache

import "testing"

func TestCache_Has(t *testing.T) {
	c := Cache{}
	c.Put("test", "123")
	if !c.Has("test") {
		t.Fatal("error")
	}
}

func TestCache_PullString(t *testing.T) {
	c := Cache{}
	c.Put("test", "123")
	if c.PullString("test") != "123" {
		t.Fatal("error")
	}
}
