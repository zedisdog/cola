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
