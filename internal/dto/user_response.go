package dto

import "github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/domain"

type UserLoginResponse struct {
	StatusCode int               `json:"statusCode"`
	Status     string            `json:"status"`
	Message    string            `json:"message"`
	Data       map[string]string `json:"data"`
}

type UserRegisterResponse struct {
	StatusCode int                 `json:"statusCode"`
	Status     string              `json:"status"`
	Message    string              `json:"message"`
	Data       UserRegisterRequest `json:"data"`
}

type UserValidateResponse struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       domain.User `json:"data"`
}

type UserGetAllDataResponse struct {
	StatusCode int           `json:"statusCode"`
	Status     string        `json:"status"`
	Message    string        `json:"message"`
	Data       []domain.User `json:"data"`
	Pagination Pagination    `json:"pagination"`
}
