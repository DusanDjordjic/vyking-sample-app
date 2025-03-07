package response

import (
	"net/http"
)

type errorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	if len(message) == 0 {
		message = http.StatusText(status)
	}

	res := errorResponse{
		StatusCode: status,
		Message:    message,
	}

	JSONResponse(w, res.StatusCode, res)
}

func NewInternalError(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusInternalServerError, "")
}

func NewBadRequest(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusBadRequest, "")
}

func NewNotFound(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusNotFound, "")
}
