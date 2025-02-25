package usecase

import (
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/domain"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/dto"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) Login(request *dto.UserLoginRequest) (string, error) {
	// Validate Request Data
	if err := pkg.ValidateStruct(request); err != nil {
		return "", apierror.APIErrorResponse{StatusCode: http.StatusBadRequest, Status: "error", Message: err}
	}

	// Call Respository
	data, err := u.repository.GetUserByUsername(request.Username)
	if err != nil {
		return "", err
	}

	// Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password)); err != nil {
		return "", apierror.APIErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Status:     "error",
			Message:    "Username or Password Wrong",
		}
	}

	// Build Token
	temporaryToken, err := pkg.GenerateJWT(data.ID, data.Name, data.Username)
	if err != nil {
		return "", err
	}

	return temporaryToken, nil
}

func (u *UserUsecase) Register(request *dto.UserRegisterRequest) error {
	// Validate Request
	if err := pkg.ValidateStruct(request); err != nil {
		return apierror.APIErrorResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err,
		}
	}

	// Check Data Username Exist or Not
	data, _ := u.repository.GetUserByUsername(request.Username)
	if data.Username != "" {
		return apierror.APIErrorResponse{
			StatusCode: http.StatusConflict,
			Status:     "error",
			Message: apierror.ValidationErrors{
				apierror.ValidationError{
					Field:      "username",
					ErrorField: "Data Username Exist, Try Other Username Data",
					Tag:        "unique",
					Value:      request.Username,
					Constraint: "unique",
				},
			},
		}
	}

	// Hashing Data
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return apierror.APIErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	// Call Repository
	if err := u.repository.CreateUser(&domain.User{
		Name:     request.Name,
		Age:      request.Age,
		Username: request.Username,
		Password: string(hashPassword),
	}); err != nil {
		return err
	}

	// Return Data
	return nil
}
