package cache

var instance = new(Cache)

func Has(key string) bool {
	return instance.Has(key)
}

func Put(key string, value interface{}) {
	instance.Put(key, value)
}

func PutWithExpire(key string, value interface{}, expire int64) {
	instance.PutWithExpire(key, value, expire)
}

func Pull(key string) interface{} {
	return instance.Pull(key)
}

func PullString(key string) string {
	return instance.PullString(key)
}

func PullInt(key string) int {
	return instance.PullInt(key)
}
