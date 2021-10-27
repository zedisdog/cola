package tools

import "testing"

func TestMul(t *testing.T) {
	a, _ := Mul("155.39", 100)
	if a != 15539 {
		t.Fatal("error")
	}
	a, _ = Mul(float64(155.39), 100)
	if a != 15539 {
		t.Fatal("error2")
	}
	a, _ = Mul(float32(155.39), 100)
	if a != 15539 {
		t.Fatal("error2")
	}
	a, _ = Mul(15539, 100)
	if a != 1553900 {
		t.Fatal("error2")
	}
}
