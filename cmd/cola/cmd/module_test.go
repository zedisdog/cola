package cmd

import (
	"fmt"
	"os"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Printf("%c", os.PathSeparator)
}
