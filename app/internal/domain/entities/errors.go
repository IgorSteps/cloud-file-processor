package entities

// baseError is our base error type.
type baseError struct {
	msg string
}

// Error implements the error interface for baseError.
func (e *baseError) Error() string {
	return e.msg
}

// InvalidInputError represents an error for invalid user input.
type InvalidInputError struct {
	baseError
}

// NewInvalidInputError creates a new instance of InvalidInputError.
func NewInvalidInputError(msg string) *InvalidInputError {
	return &InvalidInputError{
		baseError: baseError{msg: msg},
	}
}
