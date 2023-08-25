package goatcounter

import (
	"encoding/json"
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

// TestErrorResponse tests whether JSON unmarshals correctly to ErrorResponse
func TestErrorResponse(t *testing.T) {
	e := &ErrorResponse{}
	if err := json.Unmarshal(simple, e); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(complex, e); err != nil {
		t.Fatal(err)
	}
}
