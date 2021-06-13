package features

import "time"

type ITimeControl interface {
	IsTime() bool
}

func NewTimeControl(when int64) *TimeControl {
	return &TimeControl{
		when: when,
	}
}

type TimeControl struct {
	when int64
}

func (t TimeControl) IsTime() bool {
	now := time.Now().Unix()
	if t.when > now {
		return false
	}
	return true
}
