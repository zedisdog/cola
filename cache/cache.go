package cache

import (
	"sync"
	"time"
)

type item struct {
	Value  interface{}
	Expire int64
}

type cache struct {
	sync.Map
}

func (c *cache) Has(key string) (ok bool) {
	c.clear()
	_, ok = c.Load(key)
	return
}

func (c *cache) Get(key string) (v interface{}, ok bool) {
	return c.Load(key)
}

func (c *cache) Put(key string, value interface{}) {
	c.Store(key, value)
}

func (c *cache) PutWithExpire(key string, value interface{}, expire int64) {
	c.clear()
	c.Store(key, item{
		Value:  value,
		Expire: expire,
	})
}

func (c *cache) Pull(key string) (value interface{}) {
	c.clear()
	tmp, exists := c.LoadAndDelete(key)
	if tmp == nil || !exists {
		return nil
	}

	if value, ok := tmp.(item); ok {
		return value.Value
	} else {
		return tmp
	}
}

func (c *cache) PullString(key string) string {
	value := c.Pull(key)
	if value != nil {
		return value.(string)
	}
	return ""
}

func (c *cache) PullInt(key string) int {
	value := c.Pull(key)
	if value != nil {
		return value.(int)
	}
	return 0
}

func (c *cache) clear() {
	c.Range(func(key, value interface{}) bool {
		if i, ok := value.(item); ok {
			now := time.Now().Unix()
			if now >= i.Expire {
				c.Delete(key)
			}
		}
		return true
	})
}
