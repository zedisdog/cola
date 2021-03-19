package baidubce

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestClient_Handwriting(t *testing.T) {
	client := New("", "", 0)
	pwd, _ := os.Getwd()
	fp, err := os.Open(fmt.Sprintf("%s/../../test_data/3.jpg", pwd))
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()
	image, err := io.ReadAll(fp)
	res, err := client.Handwriting(HandwritingConfig{
		Image: image,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", res)
}

type A map[string]string

func (a A) Set(key string, value string) {
	a[key] = value
}

func (a A) Delete(key string) {
	delete(a, key)
}

func TestIoReader(t *testing.T) {
	a := make(A)
	a.Set("test", "test")
	fmt.Printf("%+v", a)
	a.Delete("test")
	fmt.Printf("%+v", a)
}
