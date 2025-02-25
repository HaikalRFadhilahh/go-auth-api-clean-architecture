package http

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/dto"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/usecase"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) error {
	// Create Request Dto
	var request dto.UserLoginRequest

	// Bind Request Json
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		return apierror.ErrBadRequest
	}

	// Use Login Usecase
	token, err := u.usecase.Login(&request)
	if err != nil {
		return err
	}

	// Building Response Data
	res := dto.UserLoginResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Login Success",
		Data: map[string]string{
			"token": token,
		},
	}

	// Return Json
	return pkg.HttpSuccessResponse(w, res)
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) error {
	// Variabel Request
	var request dto.UserRegisterRequest

	// Binding Request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		return apierror.ErrBadRequest
	}

	// Call Usecase Register
	if err := u.usecase.Register(&request); err != nil {
		return err
	}

	// Prepare Response Json
	res := dto.UserRegisterResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "success register user",
		Data:       request,
	}

	return pkg.HttpSuccessResponse(w, res)
}

func (u *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	return pkg.HttpSuccessResponse(w, map[string]string{
		"message": "Data All Users",
	})
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
