package tools

func Empty(data interface{}) bool {
	if data == nil {
		return true
	}

	switch data.(type) {
	case int, int32, int64, uint, uint64, uint32:
		return data == 0
	case string:
		return data == ""
	case []byte:
		return len(data.([]byte)) == 0
	}

	return true
}
