// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// UnprocessableEntityCode - A short code indicating the error code returned.
type UnprocessableEntityCode string

const (
	UnprocessableEntityCodeUnprocessableEntity UnprocessableEntityCode = "unprocessable_entity"
)

func (e UnprocessableEntityCode) ToPointer() *UnprocessableEntityCode {
	return &e
}
func (e *UnprocessableEntityCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unprocessable_entity":
		*e = UnprocessableEntityCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UnprocessableEntityCode: %v", v)
	}
}

type UnprocessableEntityError struct {
	// A short code indicating the error code returned.
	Code UnprocessableEntityCode `json:"code"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
}

func (o *UnprocessableEntityError) GetCode() UnprocessableEntityCode {
	if o == nil {
		return UnprocessableEntityCode("")
	}
	return o.Code
}

func (o *UnprocessableEntityError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *UnprocessableEntityError) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

// UnprocessableEntity - The request was well-formed but was unable to be followed due to semantic errors.
type UnprocessableEntity struct {
	Error_ UnprocessableEntityError `json:"error"`
}

var _ error = &UnprocessableEntity{}

func (e *UnprocessableEntity) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
