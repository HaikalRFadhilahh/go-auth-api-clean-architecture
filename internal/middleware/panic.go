package middleware

import (
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create Anonymous Function Exec Defer
		defer func() {
			if r := recover(); r != nil {
				_ = pkg.HttpErrorResponse(w, apierror.ErrInternalServerError)
			}
		}()

		// Forward Request
		next.ServeHTTP(w, r)
	})
}
