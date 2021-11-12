package tools

import (
	"bytes"
	"time"
)

type CTime struct {
	time.Time
}

func (c *CTime) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, "\"")
	zone := time.FixedZone("CST", int((8 * time.Hour).Seconds()))
	tmp, err := time.ParseInLocation("2006-01-02", string(b), zone)
	if err != nil {
		return
	}
	*c = CTime{tmp}
	return
}

type CJson string

func (c *CJson) UnmarshalJSON(b []byte) (err error) {
	*c = CJson(b)
	return nil
}

func (c CJson) MarshalJSON() ([]byte, error) {
	return []byte(c), nil
}
