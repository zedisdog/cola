package task

import (
	"context"
	"testing"
)

func TestQueue(t *testing.T) {
	// 三个job都对testData做修改，job1往testData添加一个1元素,job2添加2,job3添加3
	testData := make([]int, 0, 3)
	// 每个job分别设置延迟 job1 -> 3s, job2 -> 2s, job3 -> 1s
	delays := []int64{3, 2, 1}

	// 测试数据
	jobs := make([]*job, 0, 3)
	for index, delay := range delays {
		func(index int, delay int64) {
			jobs = append(jobs,
				NewJob(func(cxt context.Context) error {
					testData = append(testData, index+1) // 从1开始
					return nil
				}, nil).delay(delay),
			)
		}(index, delay)
	}

	// 开始跑
	queue := NewQueue(3, nil)
	queue.Start()
	// dispatch顺序是 job1 -> job2 -> job3
	for _, job := range jobs {
		err := queue.Dispatch(job)
		if err != nil {
			println(err.Error())
			t.Fatal("should not return err")
		}
	}

	// 等待任务跑完,Stop会等待所有任务跑完才停下.如果不关闭,可以使用time.Sleep(time.Duration(5)*time.Second)测试.
	queue.Stop()

	// 验证
	// 如果不添加延迟,应该得到的testData应该是[1,2,3], 但是因为设置了延迟,最后应该得到的testData应该是[3,2,1]
	if testData[0] != 3 || testData[1] != 2 || testData[2] != 1 {
		t.Fatal("error")
	}
}
