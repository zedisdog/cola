package tools

import (
	"regexp"
)

func IsMobile(str string) bool {
	reg := `^1[3456789]\d{9}$`
	r, _ := regexp.Compile(reg)
	return r.MatchString(str)
}

func IsTrue(str string) bool {
	if str != "" && str != "0" && str != "false" {
		return true
	}
	return false
}
