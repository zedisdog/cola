package store

import (
	"github.com/uniplaces/carbon"
	"sync"
	"time"
)

var lock sync.Mutex

type VerifyTokenInfo struct {
	Token  string
	Expire time.Time
}

type VerifyToken map[string]VerifyTokenInfo

func (v VerifyToken) Put(key string, value string) {
	lock.Lock()
	v.clear()
	v[key] = VerifyTokenInfo{
		Token:  value,
		Expire: carbon.Now().AddHours(2).SubMinutes(30).Time,
	}
	lock.Unlock()
}

func (v VerifyToken) Pull(key string) string {
	if info, ok := v[key]; ok {
		lock.Lock()
		delete(v, key)
		lock.Unlock()
		return info.Token
	}
	return ""
}

func (v VerifyToken) clear() {
	for key, value := range v {
		if carbon.Now().Lt(carbon.NewCarbon(value.Expire)) {
			delete(v, key)
		}
	}
}
