package task

import "sync"

func newSliceDriver() *sliceDriver {
	return &sliceDriver{
		jobs: make([]interface{}, 0),
	}
}

type sliceDriver struct {
	jobs []interface{}
	lock sync.Mutex
}

func (s *sliceDriver) pop() interface{} {
	if len(s.jobs) == 0 {
		return nil
	}
	s.lock.Lock()
	target := s.jobs[0]
	//tmp := make([]interface{}, 0, len(s.jobs)-1)
	//copy(tmp, s.jobs[1:])
	//s.jobs = tmp
	s.jobs = s.jobs[1:]
	s.lock.Unlock()
	return target
}

func (s *sliceDriver) put(content interface{}) {
	s.lock.Lock()
	s.jobs = append(s.jobs, content)
	s.lock.Unlock()
}
