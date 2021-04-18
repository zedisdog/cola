package task

import (
	"github.com/sirupsen/logrus"
	"github.com/zedisdog/cola/errx"
	"sync"
)

func newWorker(pool chan chan *job, wg *sync.WaitGroup, logger *logrus.Logger) *worker {
	return &worker{
		logger:     logger,
		job:        make(chan *job),
		workerPool: pool,
		wg:         wg,
		quit:       make(chan struct{}),
	}
}

type worker struct {
	logger     *logrus.Logger
	job        chan *job
	workerPool chan chan *job
	wg         *sync.WaitGroup
	quit       chan struct{}
}

func (w worker) start() {
	w.wg.Add(1)
	w.workerPool <- w.job
	go w.run()
}

func (w worker) run() {
	for j := range w.job {
		err := j.do()
		if err != nil { // todo: 如果job有设置，可以在出错时把job放回队列中
			w.log(err)
		}
		w.workerPool <- w.job
	}
	w.wg.Done()
}

func (w worker) log(err error) {
	if w.logger != nil {
		entry := w.logger.WithError(err)
		if er, ok := err.(*errx.Error); ok {
			entry = entry.WithField("stack", string(er.Stack()))
		}
		entry.Error("job is failed")
	}
}
