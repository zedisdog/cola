package tools

import (
	"github.com/sony/sonyflake"
	"sync"
)

var once sync.Once
var snow *sonyflake.Sonyflake

func GetSnow() *sonyflake.Sonyflake {
	once.Do(func() {
		snow = sonyflake.NewSonyflake(sonyflake.Settings{}) //todo: 配置
	})
	return snow
}
