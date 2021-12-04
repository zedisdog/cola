package evbus

type Event struct {
	Name   interface{}
	Params map[string]interface{}
}

func (e *Event) Param(key string) (value interface{}, ok bool) {
	if e.Params == nil {
		return nil, false
	}
	value, ok = e.Params[key]
	return
}

func (e *Event) ParamString(key string) (value string, ok bool) {
	param, ok := e.Param(key)
	if !ok {
		return
	}
	if param == nil {
		value = ""
	} else {
		value = param.(string)
	}
	return
}

func (e *Event) ParamUint64(key string) (value uint64, ok bool) {
	param, ok := e.Param(key)
	if !ok {
		return
	}
	if param == nil {
		value = 0
	} else {
		value = param.(uint64)
	}
	return
}

func (e *Event) ParamInt(key string) (value int, ok bool) {
	param, ok := e.Param(key)
	if !ok {
		return
	}
	if param == nil {
		value = 0
	} else {
		value = param.(int)
	}
	return
}
