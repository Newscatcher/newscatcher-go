// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "github.com/Newscatcher/newscatcher-go/core"
)

// Validation Error
type UnprocessableEntityError struct {
	*core.APIError
	Body *HttpValidationError
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *HttpValidationError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}