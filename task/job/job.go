package job

import (
	"github.com/zedisdog/cola/task/job/features"
	"reflect"
	"time"
)

type IJob interface {
	Handle() error
}

func WithTime(when int64) func(*job) {
	return func(j *job) {
		j.TimeControl = *features.NewTimeControl(when)
	}
}

func WithDelay(seconds int) func(*job) {
	return func(j *job) {
		j.TimeControl = *features.NewTimeControl(time.Now().Add(time.Duration(seconds) * time.Second).Unix())
	}
}

func WithParams(params ...interface{}) func(*job) {
	return func(j *job) {
		j.params = params
	}
}

// NewJob create object implements IJob
//  params:
//   callable	a closure
//   options	func(*job)
func NewJob(callable interface{}, options ...func(*job)) IJob {
	j := &job{
		callable: callable,
	}
	for _, setOption := range options {
		setOption(j)
	}
	return j
}

type job struct {
	callable interface{}
	params   []interface{}
	features.TimeControl
}

func (j job) Handle() error {
	p := make([]reflect.Value, 0, len(j.params))
	for _, item := range j.params {
		p = append(p, reflect.ValueOf(item))
	}

	values := reflect.ValueOf(j.callable).Call(p)
	if len(values) > 0 {
		for _, value := range values {
			if e, ok := value.Interface().(error); ok {
				return e
			}
		}
	}
	return nil
}
