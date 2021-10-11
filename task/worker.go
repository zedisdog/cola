package task

import (
	"github.com/zedisdog/cola/task/job"
	"sync"
)

func newWorker(pool chan chan job.IJob, wg *sync.WaitGroup) *worker {
	return &worker{
		job:        make(chan job.IJob),
		workerPool: pool,
		wg:         wg,
	}
}

type worker struct {
	job        chan job.IJob
	workerPool chan chan job.IJob
	wg         *sync.WaitGroup
}

func (w worker) start() {
	w.workerPool <- w.job
	go w.run()
}

func (w worker) run() {
	w.wg.Add(1)
	for j := range w.job {
		err := j.Handle()
		if err != nil { // todo: 如果job有设置，可以在出错时把job放回队列中
			panic(err)
		}
		w.workerPool <- w.job
	}
	w.wg.Done()
}
