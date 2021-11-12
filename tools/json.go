package tools

import (
	"bytes"
	"time"
)

type JTime = CTime

//CTime parse string time to time.Time
//Deprecated: use JTime instead
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

type JJson string

func (c *JJson) UnmarshalJSON(b []byte) (err error) {
	*c = JJson(b)
	return nil
}

func (c JJson) MarshalJSON() ([]byte, error) {
	return []byte(c), nil
}
