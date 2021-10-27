package cache

import "errors"

var Caches = map[string]*cache{
	"default": {},
}

func Instance(name ...string) (cache *cache, err error) {
	cacheName := "default"
	if len(name) > 0 {
		cacheName = name[0]
	}
	cache, ok := Caches[cacheName]
	if !ok {
		err = errors.New("cache not found")
	}
	return
}

func Has(key string) bool {
	instance, _ := Instance()
	return instance.Has(key)
}

func Put(key string, value interface{}) {
	instance, _ := Instance()
	instance.Put(key, value)
}

func PutWithExpire(key string, value interface{}, expire int64) {
	instance, _ := Instance()
	instance.PutWithExpire(key, value, expire)
}

func Pull(key string) interface{} {
	instance, _ := Instance()
	return instance.Pull(key)
}

func PullString(key string) string {
	instance, _ := Instance()
	return instance.PullString(key)
}

func PullInt(key string) int {
	instance, _ := Instance()
	return instance.PullInt(key)
}
