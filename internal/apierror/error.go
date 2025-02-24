package apierror

import "net/http"

// Struct API Error Response
type APIErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// Implementation Error Function To Error Interface
func (e *APIErrorResponse) Error() string {
	return e.Message
}

// Var Error
var (
	ErrInternalServerError = APIErrorResponse{StatusCode: http.StatusInternalServerError, Status: "error", Message: "Internal Server Error"}
	ErrPageNotFound        = APIErrorResponse{StatusCode: http.StatusNotFound, Status: "error", Message: "Not Found"}
)
