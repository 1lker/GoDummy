package errors

import "fmt"

// ErrorCode represents a unique error code
type ErrorCode string

const (
	ErrInvalidInput    ErrorCode = "INVALID_INPUT"
	ErrInvalidRange    ErrorCode = "INVALID_RANGE"
	ErrRateLimitExceeded ErrorCode = "RATE_LIMIT_EXCEEDED"
	ErrServerError     ErrorCode = "SERVER_ERROR"
	ErrTimeout        ErrorCode = "TIMEOUT"
)

// APIError represents an API error
type APIError struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError creates a new APIError
func NewError(code ErrorCode, message string, details interface{}) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// InvalidInput creates an invalid input error
func InvalidInput(message string, details interface{}) *APIError {
	return NewError(ErrInvalidInput, message, details)
}

// InvalidRange creates an invalid range error
func InvalidRange(message string, details interface{}) *APIError {
	return NewError(ErrInvalidRange, message, details)
}

// RateLimitExceeded creates a rate limit exceeded error
func RateLimitExceeded() *APIError {
	return NewError(ErrRateLimitExceeded, "Rate limit exceeded", nil)
}

// ServerError creates a server error
func ServerError(err error) *APIError {
	return NewError(ErrServerError, "Internal server error", err.Error())
}

// TimeoutError creates a timeout error
func TimeoutError() *APIError {
	return NewError(ErrTimeout, "Request timeout", nil)
}