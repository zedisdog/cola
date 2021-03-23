package task

import (
	"context"
	"errors"
	"time"
)

type jobFunc = func(cxt context.Context) error

// job 任务
type job struct {
	callable jobFunc
	params   map[string]interface{}
	when     int64
}

func NewJob(callable jobFunc, params map[string]interface{}) *job {
	return &job{
		callable: callable,
		params:   params,
	}
}

func (j job) do() (err error) {
	// todo: 利用反射把参数给到闭包里面，这样可以简化闭包的编写
	var cxt context.Context
	if j.params != nil {
		cxt = context.WithValue(context.Background(), "params", j.params)
	} else {
		cxt = context.Background()
	}

	return j.callable(cxt)
}

func (j *job) Delay(seconds int64) *job {
	j.when = time.Now().Unix() + seconds
	return j
}

func (j *job) On(t time.Time) *job {
	j.when = t.Unix()
	return j
}

// dispatcher 任务队列
type dispatcher struct {
	pool    chan *job
	link    queue
	running bool
}

func newDispatcher() *dispatcher {
	return &dispatcher{
		pool: make(chan *job),
		link: newLink(),
	}
}

func (d *dispatcher) start() {
	d.running = true
	go func() {
		for {
			if j := d.link.pop(); j != nil {
				job := j.(*job)
				// 判断是否到了执行时间
				if job.when > time.Now().Unix() {
					d.link.put(job)
					continue
				}
				d.pool <- job
			} else {
				if !d.running {
					close(d.pool)
					break
				}
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}

func (d dispatcher) put(job *job) error {
	if d.running {
		d.link.put(job)
		return nil
	} else {
		return errors.New("dispatcher is closed and will not accepts jobs")
	}
}

func (d *dispatcher) stop() {
	d.running = false
}

type queue interface {
	put(interface{})
	pop() interface{}
}
