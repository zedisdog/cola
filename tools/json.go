package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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

func (c JJson) Get(name string) (value interface{}, err error) {
	tmp := make(map[string]interface{})
	err = json.Unmarshal([]byte(c), &tmp)
	if err != nil {
		return
	}
	var ok bool
	for _, n := range strings.Split(name, ".") {
		if value == nil {
			value, ok = tmp[n]
		} else {
			value, ok = value.(map[string]interface{})[n]
		}
		if !ok {
			err = errors.New(fmt.Sprintf("value of key <%s> not found", n))
			return
		}
	}

	return
}

func (c JJson) GetString(name string) (value string, err error) {
	v, err := c.Get(name)
	if err != nil {
		return
	}
	value = v.(string)
	return
}
