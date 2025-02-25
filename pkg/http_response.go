package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
)

func HttpErrorResponse(w http.ResponseWriter, e any) error {
	// Setting Response To Json Type
	w.Header().Set("Content-type", "application/json")

	// Assert Data
	data, ok := e.(apierror.APIErrorResponse)
	// Handle Internal Server Error If Assert Error
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(apierror.ErrInternalServerError)
	}

	// If Assert Success
	w.WriteHeader(data.StatusCode)
	return json.NewEncoder(w).Encode(data)
}

func HttpSuccessResponse(w http.ResponseWriter, e any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(e)
}
