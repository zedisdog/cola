package baidubce

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestClient_Handwriting(t *testing.T) {
	client := New("", "")
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
