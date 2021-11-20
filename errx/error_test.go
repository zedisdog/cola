package errx

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	err := fmt.Errorf("read private pem file err:%s", "file not exists")
	err = Wrap(err, "321")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("......%s", err)
		}
	}()
	panic(err)
}
