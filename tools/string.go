package tools

import (
	"math/rand"
	"regexp"
	"time"
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

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
