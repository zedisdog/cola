package store

import (
	"github.com/uniplaces/carbon"
	"sync"
	"time"
)

var lock sync.Mutex

type VerifyTokenInfo struct {
	Token  string
	Expire int64
}

type VerifyToken map[string]VerifyTokenInfo

func (v VerifyToken) Has(key string) bool {
	lock.Lock()
	defer lock.Unlock()
	v.clear()
	_, ok := v[key]
	return ok
}

func (v VerifyToken) Put(key string, value string) {
	v.PutWithExpire(key, value, carbon.Now().AddHours(2).SubMinutes(30).Unix())
}

func (v VerifyToken) PutWithExpire(key string, value string, expire int64) {
	lock.Lock()
	v.clear()
	v[key] = VerifyTokenInfo{
		Token:  value,
		Expire: expire,
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
		now := time.Now().Unix()
		if now >= value.Expire {
			delete(v, key)
		}
	}
}
