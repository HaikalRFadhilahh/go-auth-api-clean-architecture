package pkg

import (
	"net/http"
	"strconv"
	"time"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/golang-jwt/jwt"
)

type JsonWebTokenDataClaims struct {
	Id       int
	Name     string
	Username string
	jwt.StandardClaims
}

func GenerateJWT(id int, name, username string) (string, error) {
	// Take Secret
	secret := []byte(GetEnv("JWT_SECRET", ""))
	expHour, err := strconv.Atoi(GetEnv("JWT_EXPIRED_HOUR", "1"))
	if err != nil {
		expHour = 1
	}

	// Create Claims Data
	dataJWT := &JsonWebTokenDataClaims{
		Id:       id,
		Name:     name,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expHour)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create Token
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, dataJWT)

	// Encrypt With Secret Data
	token, err := j.SignedString(secret)
	if err != nil {
		return "", apierror.APIErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	// Return Data
	return token, nil
}

func DecodeJWT(t string) (*JsonWebTokenDataClaims, error) {
	// Get Secret Token From ENV
	secret := []byte(GetEnv("JWT_SECRET", ""))

	// Create Claims Data
	var dataClaims JsonWebTokenDataClaims

	// Decode JWT Token
	token, err := jwt.ParseWithClaims(t, &dataClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apierror.ErrForbidden
		}

		return secret, nil
	})

	// Check Error
	if err != nil {
		return nil, apierror.ErrForbidden
	}

	// Check Valid Token
	if !token.Valid {
		return nil, apierror.ErrForbidden
	}

	// Return Data Claims
	return &dataClaims, nil
}
