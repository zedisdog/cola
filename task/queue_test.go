package task

import (
	"fmt"
	"github.com/zedisdog/cola/task/job"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	// 三个job都对testData做修改，job1往testData添加一个1元素,job2添加2,job3添加3
	testData := make([]int, 0, 3)
	// 每个job分别设置延迟 job1 -> 3s, job2 -> 2s, job3 -> 1s
	delays := []int{3, 2, 1}

	// 测试数据
	jobs := make([]job.IJob, 0, 3)
	for index, delay := range delays {
		func(index int, delay int) {
			jobs = append(jobs,
				job.NewJob(func() {
					testData = append(testData, index+1) // 从1开始
				}, job.WithDelay(delay)),
			)
		}(index, delay)
	}

	// 开始跑
	queue := NewQueue(3)
	queue.Start()
	// dispatch顺序是 job1 -> job2 -> job3
	for _, j := range jobs {
		err := queue.Dispatch(j)
		if err != nil {
			println(err.Error())
			t.Fatal("should not return err")
		}
	}

	// 等待任务跑完,Stop会等待所有任务跑完才停下.如果不关闭,可以使用time.Sleep(time.Duration(5)*time.Second)测试.
	queue.Stop()
	//time.Sleep(time.Duration(5)*time.Second)

	// 验证
	// 如果不添加延迟,应该得到的testData应该是[1,2,3], 但是因为设置了延迟,最后应该得到的testData应该是[3,2,1]
	fmt.Printf("%+v", testData)
	if testData[0] != 3 || testData[1] != 2 || testData[2] != 1 {
		t.Fatal("error")
	}
}

func TestNormal(t *testing.T) {
	queue := NewQueue(1)
	queue.Start()
	_ = queue.Dispatch(job.NewJob(func() {
		println("执行任务来了3")
	}, job.WithDelay(3)))

	_ = queue.Dispatch(job.NewJob(func() {
		println("执行任务来了2")
	}, job.WithDelay(2)))

	_ = queue.Dispatch(job.NewJob(func() {
		println("执行任务来了20")
	}, job.WithDelay(20)))
	queue.Stop()
}

func TestMuch(t *testing.T) {
	result := 0
	queue := NewQueue(10)
	queue.Start()
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		//randPush := rand.Intn(5)
		//println("randpush", randPush)
		//time.Sleep(time.Duration(randPush) * time.Second)
		func(i int) {
			_ = queue.Dispatch(job.NewJob(func() {
				lock.Lock()
				result++
				lock.Unlock()
			}) /*.Delay(int64(randPush))*/)
		}(i)
	}
	queue.Stop()
	println(result)
}
