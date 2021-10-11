package task

import (
	"github.com/zedisdog/cola/task/job"
	"testing"
)

func TestRing(t *testing.T) {
	r := &ring{
		size:   MIN,
		buffer: make([]job.IJob, MIN),
	}

	if r.pop() != nil {
		t.Fatal("should pop nil")
	}
	if r.size != MIN {
		t.Fatal("should equal MIN")
	}
	if r.r != 0 {
		t.Fatal("should equal 0")
	}
	if r.w != 0 {
		t.Fatal("should equal 0")
	}

	r.put(nil)
	if r.size != MIN {
		t.Fatal("should equal MIN")
	}
	if r.r != 0 {
		t.Fatal("should equal 0")
	}
	if r.w != 1 {
		t.Fatal("should equal 1")
	}

	for i := 0; i < 10; i++ {
		r.put(nil)
	}
	if r.size != MIN*2 {
		t.Fatal("should equal MIN*2")
	}
	if r.r != 0 {
		t.Fatal("should equal 0")
	}
	if r.w != 11 {
		t.Fatal("should equal 11")
	}

	for i := 0; i < 11; i++ {
		r.pop()
	}
	if r.size != MIN {
		t.Fatal("should equal MIN")
	}
	if r.r != 0 {
		t.Fatal("should equal 0")
	}
	if r.w != 0 {
		t.Fatal("should equal 0")
	}
}

func TestRingNext(t *testing.T) {
	r := &ring{
		size:   MIN,
		buffer: make([]job.IJob, MIN),
	}

	r.put(nil)

	r.next()

	if r.w != 2 {
		t.Fatal("w should be 2")
	}
	if r.r != 1 {
		t.Fatal("r should be 1")
	}

	for i := 0; i < MIN; i++ {
		r.next()
	}

	if r.w != 2 {
		t.Fatal("w should be 2")
	}
	if r.r != 1 {
		t.Fatal("r should be 1")
	}
}
