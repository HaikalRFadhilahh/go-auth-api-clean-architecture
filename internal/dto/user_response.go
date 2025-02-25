package dto

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
