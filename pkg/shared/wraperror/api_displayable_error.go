package wraperror

import "errors"

type ApiDisplayableError struct {
	code       int
	httpStatus int
	message    interface{}
	err        error
}

func NewApiDisplayableError(
	code int,
	httpStatus int,
	message interface{},
	err string,
) *ApiDisplayableError {
	return &ApiDisplayableError{
		code:       code,
		httpStatus: httpStatus,
		message:    message,
		err:        errors.New(err),
	}
}

func (err *ApiDisplayableError) Error() string {
	if err.err != nil {
		return err.err.Error()
	}
	if message, messageIsString := err.message.(string); messageIsString {
		return message
	}

	return "Unknown error"
}

func (err *ApiDisplayableError) Unwrap() error {
	return err.err
}

func (err *ApiDisplayableError) Message() interface{} {
	return err.message
}

func (err *ApiDisplayableError) HttpStatus() int {
	return err.httpStatus
}

func (err *ApiDisplayableError) ErrorCode() int {
	return err.code
}
