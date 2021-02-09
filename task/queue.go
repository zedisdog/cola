package task

import (
	"github.com/sirupsen/logrus"
	"sync"
)

func NewQueue(max int, logger *logrus.Logger) *Queue {
	if max < 1 {
		panic("worker count can not little than 1")
	}
	return &Queue{
		dispatcher: newDispatcher(),
		wg:         &sync.WaitGroup{},
		logger:     logger,
		workerPool: make(chan chan *job, max),
		maxWorker:  max,
	}
}

type Queue struct {
	dispatcher *dispatcher
	wg         *sync.WaitGroup
	logger     *logrus.Logger
	workerPool chan chan *job
	maxWorker  int
}

func (q *Queue) Start() {
	for i := 0; i < q.maxWorker; i++ {
		worker := newWorker(q.workerPool, q.wg, q.logger)
		worker.start()
	}
	q.dispatcher.start()
	go q.run()
}

func (q Queue) run() {
	for job := range q.dispatcher.pool {
		worker := <-q.workerPool
		worker <- job
	}

	for worker := range q.workerPool {
		close(worker)
	}
}

func (q *Queue) Stop() {
	q.log("queue is shutting down...")
	q.dispatcher.stop()
	q.wg.Wait()
	close(q.workerPool)
	q.log("queue is shutdown")
}

func (q Queue) log(args ...interface{}) {
	if q.logger != nil {
		q.logger.Info(args...)
	}
}

func (q Queue) Dispatch(job *job) error {
	return q.dispatcher.put(job)
}
