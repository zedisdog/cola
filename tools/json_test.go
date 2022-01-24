package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestCTime_UnmarshalJSON(t *testing.T) {
	type a struct {
		A *JTime `json:"a"`
	}
	var aa a
	j := `{"a": "2006-02-06"}`
	err := json.Unmarshal([]byte(j), &aa)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJJson_UnmarshalJSON(t *testing.T) {
	type a struct {
		A JJson `json:"a"`
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

func TestJJson_MarshalJSON(t *testing.T) {
	type a struct {
		A JJson `json:"a"`
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

func TestJJson_Get(t *testing.T) {
	type a struct {
		A JJson `json:"a"`
	}

	aa := a{
		A: `{"b": {"c":{"d": "e"}}}`,
	}

	value, err := aa.A.Get("b.c")
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := value.(map[string]interface{}); !ok {
		t.Fatal("error")
	} else {
		if v["d"].(string) != "e" {
			t.Fatal("error")
		}
	}

	value, err = aa.A.Get("b.c.d")
	if err != nil {
		t.Fatal(err)
	}

	if value != interface{}("e") {
		t.Fatal("error")
	}

	value, err = aa.A.GetString("b.c.d")
	if err != nil {
		t.Fatal(err)
	}

	if value != "e" {
		t.Fatal("error")
	}
}
