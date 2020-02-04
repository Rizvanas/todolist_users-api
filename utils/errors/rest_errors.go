package errors

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

// NewBadRequestError returns bad request error with specified message
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: 400,
		Error:      "bad_request",
	}
}

// NewNotFoundError returns not found error with specified message
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: 404,
		Error:      "not_found",
	}
}
