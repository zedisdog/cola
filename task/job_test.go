package task

import (
	"context"
	"testing"
)

func TestQueue_Dispatcher(t *testing.T) {
	d := newDispatcher()
	err := d.put(nil)
	if err == nil {
		t.Fatal("should return err when dispatcher is not running")
	}

	d.start()
	jobs := make([]*job, 0, 3)
	delays := []int64{3, 2, 1}
	testData := make([]int, 0, 3)
	for index, delay := range delays {
		jobs = append(jobs,
			NewJob(func(cxt context.Context) error {
				testData = append(testData, index+1) // 从1开始
				return nil
			}, nil).Delay(delay),
		)
	}
	for _, job := range jobs {
		err = d.put(job)
		if err != nil {
			t.Fatal("should not return err")
		}
	}
	d.stop()

	for j := range d.pool {
		err := j.do()
		if err != nil {
			t.Fatal("should not return err")
		}
	}

	if testData[0] != 3 || testData[1] != 2 || testData[2] != 1 {
		t.Fatal("error")
	}
}
