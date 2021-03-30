package store

import (
	"github.com/uniplaces/carbon"
	"testing"
	"time"
)

func TestVerifyToken_Expire(t *testing.T) {
	s := &VerifyToken{}
	s.PutWithExpire("test", "hehehe", carbon.Now().AddSeconds(1).Unix())
	if !s.Has("test") {
		t.Fatal("error1")
	}
	time.Sleep(1 * time.Second)
	if s.Has("test") {
		t.Fatal("error2")
	}
}
