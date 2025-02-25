package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
)

func CheckJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Token Authorization
		tokenJWT := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		// Check JWT
		claims, err := pkg.DecodeJWT(tokenJWT)
		if err != nil {
			pkg.HttpErrorResponse(w, err)
			return
		}

		// Save The Claims
		jwtContext := context.WithValue(r.Context(), "jwtClaims", claims)

		// Next Request is Valid
		next.ServeHTTP(w, r.WithContext(jwtContext))
	})
}
