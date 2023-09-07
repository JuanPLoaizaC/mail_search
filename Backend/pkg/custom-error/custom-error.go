package customerror

type CustomError struct {
	ErrorText string `json:"error,omitempty"`
}

func (e *CustomError) Error() string {
	return e.ErrorText
}
