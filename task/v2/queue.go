package v2

import (
	"errors"
	"github.com/zedisdog/cola/task/v2/job"
	"sync"
)

func NewQueue(max int, opts ...func(q *Queue)) *Queue {
	if max < 1 {
		panic("worker count can not little than 1")
	}
	q := &Queue{
		workerPool: make(chan chan job.IJob, max),
		maxWorker:  max,
	}
	for _, opt := range opts {
		opt(q)
	}
	q.dispatcher = NewDispatcher()
	return q
}

type Queue struct {
	dispatcher *dispatcher
	running    bool
	workerPool chan chan job.IJob
	maxWorker  int
	wg         sync.WaitGroup
}

func (q *Queue) Start() {
	q.running = true
	for i := 0; i < q.maxWorker; i++ {
		worker := newWorker(q.workerPool, &q.wg)
		worker.start()
	}
	go q.dispatcher.run()
	go q.run()
}

func (q *Queue) run() {
	for j := range q.dispatcher.out {
		worker := <-q.workerPool
		worker <- j
	}

	for worker := range q.workerPool {
		close(worker)
	}
}

func (q *Queue) Stop() {
	q.running = false
	close(q.dispatcher.in)
	q.wg.Wait()
	close(q.workerPool)
}

func (q *Queue) Dispatch(job job.IJob) error {
	if q.running {
		q.dispatcher.in <- job
		return nil
	} else {
		return errors.New("queue is shutdown")
	}
}
