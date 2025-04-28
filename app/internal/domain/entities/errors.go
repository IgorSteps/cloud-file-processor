package entities

type InvalidError struct {
	Message string
}

func NewInvalidError(msg string) *InvalidError {
	return &InvalidError{
		Message: msg,
	}
}

func (e *InvalidError) Error() string {
	return e.Message
}
