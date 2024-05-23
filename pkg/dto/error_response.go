package dto

// BaseResponse struct
type BaseErrorResponse struct {
	Code    int            `json:"code"`
	Message interface{}    `json:"message"`
	Error   *ErrorResponse `json:"error"`
}

// ErrorResponse struct
type ErrorResponse struct {
	Details interface{} `json:"details"`
}
