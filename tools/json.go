package tools

import (
	"bytes"
	"time"
)

type cTime struct {
	time.Time
}

func (c *cTime) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, "\"")
	zone := time.FixedZone("CST", int((8 * time.Hour).Seconds()))
	tmp, err := time.ParseInLocation("2006-01-02", string(b), zone)
	if err != nil {
		return
	}
	*c = cTime{tmp}
	return
}
