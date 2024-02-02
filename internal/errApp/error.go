package errapp

import (
	"encoding/json"
	"fmt"
)

type ErrorFields map[string]string
type ErrorParams map[string]string

var (
	ErrNotFound = NewErrApp(nil, "not found", "", "S-00059")
)

type ErrApp struct {
	Err              error       `json:"-"`
	Message          string      `json:"message,omitempty"`
	DeveloperMessage string      `json:"developer_message,omitempty"`
	Code             string      `json:"code,omitempty"`
	Fields           ErrorFields `json:"fields,omitempty"`
	Params           ErrorParams `json:"params,omitempty"`
}

func NewErrApp(err error, message, developerMessage, code string) *ErrApp {
	return &ErrApp{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}
func (e *ErrApp) WithFields(fields ErrorFields) {
	e.Fields = fields
}

func (e *ErrApp) WithParams(params ErrorParams) {
	e.Params = params
}

// TODO: доделаем
func (e *ErrApp) Error() string {
	return e.Message
}

func (e *ErrApp) Unwrap() error {
	return e.Err
}

func (e *ErrApp) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func BadRequestError(message, devMessage string) *ErrApp {
	return NewErrApp(fmt.Errorf(message), message, devMessage, "RT-00001")
}

func systemErr(err error) *ErrApp {
	return NewErrApp(err, "internal system error", err.Error(), "S-00023")
}
