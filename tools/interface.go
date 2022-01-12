package tools

func Empty(data interface{}) bool {
	if data == nil {
		return true
	}

	switch data.(type) {
	case int:
		return data.(int) == 0
	case int32:
		return data.(int32) == 0
	case int64:
		return data.(int64) == 0
	case uint:
		return data.(uint) == 0
	case uint64:
		return data.(uint64) == 0
	case uint32:
		return data.(uint32) == 0
	case string:
		return data.(string) == ""
	case []byte:
		return len(data.([]byte)) == 0
	}

	return true
}
