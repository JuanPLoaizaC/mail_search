package customerror

// Definición del tipo CustomError
type CustomDbError struct {
	ErrorText string `json:"error,omitempty"`
}

// Implementación del método Error() para cumplir con la interfaz error
func (e *CustomDbError) Error() string {
	return e.ErrorText
}

// type CustomDbError struct {
// 	ErrorText string `json:"error,omitempty"`
// }

// func NewCustomDbError(statusCode int, errMessage string) *CustomDbError {
// 	customErr := &CustomDbError{
// 		ErrorText: errMessage,
// 	}
// 	return customErr
// }
