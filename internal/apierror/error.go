package apierror

import (
	"net/http"
)

// Struct API Error Response
type APIErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    any    `json:"message"`
}

// Implementation Error Function To Error Interface
func (e APIErrorResponse) Error() string {
	return ""
}

type ValidationError struct {
	Field      string `json:"field"`
	ErrorField string `json:"error"`
	Tag        string `json:"tag"`
	Value      any    `json:"value"`
	Constraint string `json:"constraint"`
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return ""
}

// Var Error
var (
	ErrInternalServerError = APIErrorResponse{StatusCode: http.StatusInternalServerError, Status: "error", Message: "Internal Server Error"}
	ErrPageNotFound        = APIErrorResponse{StatusCode: http.StatusNotFound, Status: "error", Message: "Not Found"}
	ErrBadRequest          = APIErrorResponse{StatusCode: http.StatusBadRequest, Status: "error", Message: "Bad Request"}
)
