package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestCTime_UnmarshalJSON(t *testing.T) {
	type a struct {
		A *CTime `json:"a"`
	}
	var aa a
	j := `{"a": "2006-02-06"}`
	err := json.Unmarshal([]byte(j), &aa)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCJson_UnmarshalJSON(t *testing.T) {
	type a struct {
		A CJson `json:"a"`
	}

	var aa a
	j := `{"a": {"b": "c"}}`
	err := json.Unmarshal([]byte(j), &aa)
	if err != nil {
		t.Fatal(err)
	}
	if aa.A != "{\"b\": \"c\"}" {
		t.Fatal(errors.New("should be {\"b\": \"c\"}"))
	}
}

func TestCJson_MarshalJSON(t *testing.T) {
	type a struct {
		A CJson `json:"a"`
	}

	aa := a{
		A: "{\"b\": \"c\"}",
	}

	bytes, err := json.Marshal(aa)
	if err != nil {
		t.Fatal(err)
	}

	if string(bytes) != "{\"a\":{\"b\":\"c\"}}" {
		t.Fatal(errors.New(fmt.Sprintf("actually: %s", string(bytes))))
	}
}
