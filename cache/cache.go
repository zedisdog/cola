package cache

import (
	"github.com/uniplaces/carbon"
	"sync"
	"time"
)

var lock sync.Mutex

type item struct {
	Value  interface{}
	Expire int64
}

type Cache map[string]item

func (v Cache) Has(key string) bool {
	lock.Lock()
	defer lock.Unlock()
	v.clear()
	_, ok := v[key]
	return ok
}

func (v Cache) Put(key string, value interface{}) {
	v.PutWithExpire(key, value, carbon.Now().AddHours(2).SubMinutes(30).Unix())
}

func (v Cache) PutWithExpire(key string, value interface{}, expire int64) {
	lock.Lock()
	v.clear()
	v[key] = item{
		Value:  value,
		Expire: expire,
	}
	lock.Unlock()
}

func (v Cache) Pull(key string) interface{} {
	if item, ok := v[key]; ok {
		lock.Lock()
		delete(v, key)
		lock.Unlock()
		return item.Value
	}
	return nil
}

func (v Cache) PullString(key string) string {
	value := v.Pull(key)
	if value != nil {
		return value.(string)
	}
	return ""
}

func (v Cache) PullInt(key string) int {
	value := v.Pull(key)
	if value != nil {
		return value.(int)
	}
	return 0
}

func (v Cache) clear() {
	for key, value := range v {
		now := time.Now().Unix()
		if now >= value.Expire {
			delete(v, key)
		}
	}
}
