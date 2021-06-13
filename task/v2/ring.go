package v2

import (
	"errors"
	"github.com/zedisdog/cola/task/v2/job"
)

const MIN = 10

type ring struct {
	size   int
	w      int
	r      int
	buffer []job.IJob
}

func (r *ring) put(job job.IJob) {
	r.buffer[r.w] = job
	r.w++
	if r.w == r.size {
		r.w = 0
	}
	if r.r == r.w {
		r.grow()
	}
}

func (r *ring) geek() job.IJob {
	if r.r == r.w { // Empty
		panic(errors.New("empty"))
	}
	return r.buffer[r.r]
}

func (r *ring) next() {
	r.put(r.pop())
}

func (r *ring) pop() job.IJob {
	if r.r == r.w {
		return nil
	}

	j := r.buffer[r.r]
	r.r++

	if r.r == r.size {
		r.r = 0
	}
	if r.r == r.w {
		r.reduce()
	}
	return j
}

func (r *ring) grow() {
	var size int
	if r.size < 1024 {
		size = r.size * 2
	} else {
		size = r.size + r.size/4
	}

	buf := make([]job.IJob, size)

	copy(buf[0:], r.buffer[r.r:])
	copy(buf[r.size-r.r:], r.buffer[0:r.r])

	r.r = 0
	r.w = r.size
	r.size = size
	r.buffer = buf
}

func (r *ring) reduce() {
	if r.size/2 >= MIN {
		size := r.size / 2
		buf := make([]job.IJob, size)
		r.r = 0
		r.w = 0
		r.size = size
		r.buffer = buf
	}
}

func (r *ring) IsEmpty() bool {
	return r.len() == 0
}

func (r *ring) len() int {
	if r.r == r.w {
		return 0
	}

	if r.w > r.r {
		return r.w - r.r
	}
	return r.size - r.r + r.w
}
