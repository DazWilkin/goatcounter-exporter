package goatcounter

import "encoding/json"

type ErrorResponse struct {
	XError string `json:"error,omitempty"`
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
