package tools

import (
	"github.com/sony/sonyflake"
	"sync"
	"time"
)

var once sync.Once
var snow *sonyflake.Sonyflake

func WithStartTime(t time.Time) func(*sonyflake.Settings) {
	return func(settings *sonyflake.Settings) {
		settings.StartTime = t
	}
}

func WithMachineIDFunc(f func() (uint16, error)) func(*sonyflake.Settings) {
	return func(settings *sonyflake.Settings) {
		settings.MachineID = f
	}
}

func WithCheckMachineIDFunc(f func(uint16) bool) func(*sonyflake.Settings) {
	return func(settings *sonyflake.Settings) {
		settings.CheckMachineID = f
	}
}

func GetSnow(setters ...func(*sonyflake.Settings)) *sonyflake.Sonyflake {
	setting := sonyflake.Settings{}
	for _, setter := range setters {
		setter(&setting)
	}
	once.Do(func() {
		snow = sonyflake.NewSonyflake(setting)
	})
	return snow
}
