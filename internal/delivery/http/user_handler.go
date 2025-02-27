package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/dto"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/usecase"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
	"github.com/gorilla/mux"
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

func (u *UserHandler) Validate(w http.ResponseWriter, r *http.Request) error {
	// Ekstract Token From Header Autorization
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

	// Call Usecase
	datas, err := u.usecase.Validate(token)
	if err != nil {
		return err
	}

	// Building Json Response Struct
	res := dto.UserValidateResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Data Users Decode From Authorization JWT Token",
		Data:       datas,
	}

	// Return Success Response
	return pkg.HttpSuccessResponse(w, res)
}

func (u *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	// Take Data Query Data Header
	search := r.URL.Query().Get("search")
	activePage, err := strconv.Atoi(r.URL.Query().Get("activePage"))
	if err != nil {
		activePage = 1
	}

	// Exec Usecase Get All Data User
	datas, pagination, err := u.usecase.GetDataUser(search, activePage)
	if err != nil {
		return err
	}

	// Building Response
	res := dto.UserGetAllDataResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Data All Users",
		Data:       datas,
		Pagination: pagination,
	}

	// Return Response
	return pkg.HttpSuccessResponse(w, res)
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	// Get Id Data
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return apierror.ErrBadRequest
	}

	// Get Data Request
	var request dto.UserUpdateRequest

	// Decode Request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		return apierror.ErrBadRequest
	}
	request.Id = id

	// Use Usecase Update User
	data, err := u.usecase.UpdateUser(&request)
	if err != nil {
		return err
	}

	// Building Response Structute
	res := dto.UserUpdateResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "success Updated Data",
		Data:       data,
	}

	// Return Response To Users
	return pkg.HttpSuccessResponse(w, res)
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	// Take Params
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return apierror.ErrBadRequest
	}

	// Call usecase Function
	data, err := u.usecase.DeleteUser(id)
	if err != nil {
		return err
	}

	// Building The Response
	res := dto.UserDeleteResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    fmt.Sprintf("Success Delete Data With ID {%v}", id),
		Data:       data,
	}

	return pkg.HttpSuccessResponse(w, res)
}
