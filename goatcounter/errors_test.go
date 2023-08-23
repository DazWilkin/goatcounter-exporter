package goatcounter

import (
	"encoding/json"
	"log"
	"testing"
)

var (
	simple = []byte(`{
		"error": "oh noes!"
	}`)
	complex = []byte(`{
		"errors": {
			"key":     ["error1", "error2"],
			"another": ["oh noes!"]
		}
	}`)
)

func TestError(t *testing.T) {
	e := &ErrorResponse{}
	if err := json.Unmarshal(simple, e); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(complex, e); err != nil {
		log.Fatal(err)
	}
}
