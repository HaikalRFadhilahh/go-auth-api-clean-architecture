package http

import (
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func (u *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
