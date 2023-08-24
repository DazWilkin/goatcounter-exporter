package goatcounter

import (
	"encoding/json"
)

var _ error = (*ErrorResponse)(nil)

// ErrorResponse is a type that represents GoatCounter API error responses
// ErrorResponse implements the error interface
type ErrorResponse struct {
	// GError is used to avoid conflicting with the interface's Error method
	// The field must be Exported for JSON marshaling
	GError string `json:"error,omitempty"`
	Errors Errors `json:"errors,omitempty"`
}

func (e ErrorResponse) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(b)
}

type Errors map[string][]string
