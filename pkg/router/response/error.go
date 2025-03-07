package response

import (
	"net/http"
)

type errorResponse struct {
	StatusCode int      `json:"-"`
	Errors     []string `json:"errors"`
}

func ErrorResponse(w http.ResponseWriter, status int, errs ...string) {
	if len(errs) == 0 {
		errs = append(errs, http.StatusText(status))
	}

	res := errorResponse{
		StatusCode: status,
		Errors:     errs,
	}

	JSONResponse(w, res.StatusCode, res)
}

func NewInternalError(w http.ResponseWriter, errs ...string) {
	ErrorResponse(w, http.StatusInternalServerError, errs...)
}

func NewBadRequest(w http.ResponseWriter, errs ...string) {
	ErrorResponse(w, http.StatusBadRequest, errs...)
}

func NewNotFound(w http.ResponseWriter, errs ...string) {
	ErrorResponse(w, http.StatusNotFound, errs...)
}
