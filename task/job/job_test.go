package job

import (
	"errors"
	"github.com/zedisdog/cola/task/job/features"
	"testing"
	"time"
)

type A struct {
	paramA int
	paramB string
	*features.TimeControl
}

func (a A) Handle() error {
	println(123)
	return nil
}

func TestNormal(t *testing.T) {
	var a interface{} = &A{
		paramA: 1,
		paramB: "123",
	}

	_, ok := a.(IJob)
	if ok == false {
		t.Fatal("should implements IJob")
	}
	_, ok = a.(features.ITimeControl)
	if ok == false {
		t.Fatal("should implements ITimeControl")
	}
}

func TestNewJob(t *testing.T) {
	j := NewJob(func() {})
	if e := j.Handle(); e != nil {
		t.Fatal(e)
	}

	e := errors.New("123")
	j = NewJob(func() error { return e })
	if er := j.Handle(); er != e {
		t.Fatal("should equal e")
	}

	b := new(int)
	*b = 1
	j = NewJob(func(a int) { *b = a + 1 }, WithParams(1))
	if e := j.Handle(); e != nil {
		t.Fatal(e)
	}

	j = NewJob(func() {}, WithTime(time.Now().Add(1*time.Second).Unix()))
	if jj, ok := j.(features.ITimeControl); !ok {
		t.Fatal("should implements ITimeControl")
	} else {
		if jj.IsTime() == true {
			t.Fatal("should be false")
		}
		time.Sleep(1 * time.Second)
		if jj.IsTime() == false {
			t.Fatal("should be true")
		}
	}
}
