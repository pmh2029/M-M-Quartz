package wraperror

import "net/http"

type ValidationError struct {
	code     int
	messages map[string]interface{}
	Err      error
}

func NewValidationError(
	code int,
	messages map[string]interface{},
	err string,
) *ValidationError {
	return &ValidationError{
		code:     code,
		messages: messages,
		Err: NewApiDisplayableError(
			code,
			http.StatusOK,
			messages,
			err,
		),
	}
}

func (err *ValidationError) Error() string {
	return err.Err.Error()
}

func (err *ValidationError) Unwrap() error {
	return err.Err
}

func (err *ValidationError) ErrorCode() int {
	return err.code
}
