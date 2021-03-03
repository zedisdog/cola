package tools

import "regexp"

func IsMobile(str string) bool {
	reg := `^1[3456789]\d{9}$`
	r, _ := regexp.Compile(reg)
	return r.MatchString(str)
}
