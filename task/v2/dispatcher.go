package v2

import (
	"github.com/zedisdog/cola/task/v2/job"
	"github.com/zedisdog/cola/task/v2/job/features"
)

func NewDispatcher() *dispatcher {
	return &dispatcher{
		out: make(chan job.IJob),
		in:  make(chan job.IJob),
		buf: &ring{
			size:   MIN,
			buffer: make([]job.IJob, MIN),
		},
	}
}

type dispatcher struct {
	out     chan job.IJob
	in      chan job.IJob
	buf     *ring
	running bool
}

func (d *dispatcher) run() {
	for {
		j, ok := <-d.in

		if !ok {
			close(d.out)
			return
		}

		if d.canRun(j) {
			select {
			case d.out <- j:
				continue
			default:
			}
		}

		d.buf.put(j)
		for !d.buf.IsEmpty() {
			j := d.geek()

			if j != nil {
				select {
				case d.out <- j:
					d.buf.pop()
				default:
				}
			}

			select {
			case j, ok := <-d.in:
				if ok {
					d.buf.put(j)
				}
			default:
			}
		}
	}
}

func (d dispatcher) geek() job.IJob {
	for i := 0; i < d.buf.len(); i++ {
		j := d.buf.geek()
		if d.canRun(j) {
			return j
		}
		d.buf.next()
	}
	return nil
}

func (d dispatcher) canRun(j job.IJob) bool {
	timeControl, ok := j.(features.ITimeControl)
	if (ok && timeControl.IsTime()) || !ok {
		return true
	}
	return false
}
