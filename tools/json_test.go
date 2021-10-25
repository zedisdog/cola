package tools

import (
	"encoding/json"
	"testing"
)

func TestCTime_UnmarshalJSON(t *testing.T) {
	type a struct {
		A *cTime `json:"a"`
	}
	var aa a
	j := `{"a": "2006-02-06"}`
	err := json.Unmarshal([]byte(j), &aa)
	if err != nil {
		t.Fatal(err)
	}
}
