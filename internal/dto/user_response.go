package dto

import "github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/domain"

type GetUserResponse struct {
	StatusCode int           `json:"statusCode"`
	Status     string        `json:"status"`
	Message    string        `json:"message"`
	Data       []domain.User `json:"data"`
}

type GetUserById struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       domain.User `json:"data"`
}
