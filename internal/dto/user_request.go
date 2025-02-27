package dto

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Age      int    `json:"age" validate:"required,number,min=10"`
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=3,max=255"`
}

type UserUpdateRequest struct {
	Id       int    `json:"-" validate:"min=1"`
	Name     string `json:"name" validate:"omitempty,min=3,max=255"`
	Age      int    `json:"age" validate:"omitempty,min=10"`
	Username string `json:"username" validate:"omitempty,min=3,max=255"`
}
