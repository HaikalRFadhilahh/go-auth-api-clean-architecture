package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
)

func HttpErrorResponse(w http.ResponseWriter, e any) {
	// Setting Response To Json Type
	w.Header().Set("Content-type", "application/json")

	// Assert Data
	data, ok := e.(apierror.APIErrorResponse)
	// Handle Internal Server Error If Assert Error
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apierror.ErrInternalServerError)
		return
	}

	// If Assert Success
	w.WriteHeader(data.StatusCode)
	json.NewEncoder(w).Encode(data)
}

func HttpSuccessResponse(w http.ResponseWriter, e any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(e)
}
