package customerror

import (
	"net/http"

	"github.com/go-chi/render"
)

// CustomError it is a struc that will be used when you want to handle
// an error in a personalized way.
type CustomHttpError struct {
	Status    int    `json:"status"`
	ErrorText string `json:"error,omitempty"`
}

// NewCustomError works as the conntrucutor of the CustomError struc
func NewCustomHttpError(statusCode int, errMessage string) *CustomHttpError {
	customErr := &CustomHttpError{
		Status:    statusCode,
		ErrorText: errMessage,
	}
	return customErr
}

// ErrorResponseHandling function that is implemented when responding
// to an http request with an error
func (cr *CustomHttpError) ErrorResponseHandling(w http.ResponseWriter, r *http.Request) {
	render.Status(r, cr.Status)
	render.JSON(w, r, cr)
}
