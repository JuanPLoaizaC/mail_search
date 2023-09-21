package customerror

// CustomError allows adding a custom message to an error
type CustomError struct {
	ErrorText string `json:"error,omitempty"`
}

// Error implements the error interface and returns the error text
func (e *CustomError) Error() string {
	return e.ErrorText
}
