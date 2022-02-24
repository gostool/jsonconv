package jsonconv

import (
	"encoding/json"
	"testing"
)

func TestUnderline2Camel(t *testing.T) {
	resp := map[string]string{
		"a": "hello",
		"a_B": "hello",
	}
	res, err := json.Marshal(JsonCamelCase{Value: resp})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("res:%s\n", res)
}

func TestCamel2Underline(t *testing.T) {
	resp := map[string]string{
		"a": "hello",
		"aB": "hell_o",
		"aBc": "aB",
	}
	res, err := json.Marshal(JsonSnakeCase{Value: resp})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("res:%s\n", res)
}
