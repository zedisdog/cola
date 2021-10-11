package task

import (
	"errors"
	"github.com/zedisdog/cola/errx"
	"github.com/zedisdog/cola/task/job"
	"sync"
)

var queues = make(map[string]*queue)

func Queue(name ...string) *queue {
	queueName := "default"
	if len(name) > 0 && name[0] != "" {
		queueName = name[0]
	}
	return getOrCreate(queueName)
}

func getOrCreate(name string) *queue {
	if q, ok := queues[name]; ok {
		return q
	}

	queues[name] = NewQueue()
	return queues[name]
}

func WithWorkNum(num int) func(*queue) {
	if num < 1 {
		panic(errx.New("worker count can not little than 1"))
	}
	return func(q *queue) {
		q.maxWorker = num
	}
}

func IsRunning() bool {
	for _, queue := range queues {
		if queue.Running {
			return true
		}
	}

	return false
}

func Stop() {
	for _, queue := range queues {
		if queue.Running {
			queue.Stop()
		}
	}
}

func NewQueue(opts ...func(q *queue)) *queue {
	q := &queue{
		maxWorker: 1,
	}
	for _, opt := range opts {
		opt(q)
	}
	q.dispatcher = NewDispatcher()
	return q
}

type queue struct {
	dispatcher *dispatcher
	Running    bool
	workerPool chan chan job.IJob
	maxWorker  int
	wg         sync.WaitGroup
}

func (q *queue) Start() {
	if q.workerPool == nil {
		q.workerPool = make(chan chan job.IJob, q.maxWorker)
	}
	q.Running = true
	for i := 0; i < q.maxWorker; i++ {
		worker := newWorker(q.workerPool, &q.wg)
		worker.start()
	}
	go q.dispatcher.run()
	go q.run()
}

func (q *queue) run() {
	for j := range q.dispatcher.out {
		worker := <-q.workerPool
		worker <- j
	}

	for worker := range q.workerPool {
		close(worker)
	}
}

func (q *queue) Stop() {
	q.Running = false
	close(q.dispatcher.in)
	q.wg.Wait()
	close(q.workerPool)
}

func (q *queue) Dispatch(job job.IJob) error {
	if q.Running {
		q.dispatcher.in <- job
		return nil
	} else {
		return errors.New("queue is shutdown")
	}
}
