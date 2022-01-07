package funcs

import (
	"fmt"
	"strings"
)

func ParamList(l interface{}) string {
	var s []string
	if ml, ok := l.(map[string]string); ok {
		s = make([]string, 0, len(ml))
		for k, v := range ml {
			s = append(s, fmt.Sprintf("%s %s", k, v))
		}
	} else if sl, ok := l.([]string); ok {
		s = sl
	}
	return strings.Join(s, ", ")
}

func ReturnList(l interface{}) string {
	if ml, ok := l.(map[string]string); ok {
		var s []string
		for k, v := range ml {
			s = append(s, fmt.Sprintf("%s %s", k, v))
		}
		return fmt.Sprintf("(%s)", strings.Join(s, ", "))
	} else if sl, ok := l.([]string); ok {
		if len(sl) > 1 {
			return fmt.Sprintf("(%s)", strings.Join(sl, ", "))
		} else {
			return sl[0]
		}
	}
	return ""
}
