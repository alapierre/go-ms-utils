package commons

import (
	"encoding/json"
	"errors"
	"io"
)

type RestError struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	StatusCode       int
}

func MakeError(body io.ReadCloser) error {
	var restError RestError

	err := json.NewDecoder(body).Decode(&restError)
	if err != nil {
		return err
	}
	return errors.New(restError.ErrorDescription)
}
