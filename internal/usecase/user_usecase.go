package usecase

import (
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/domain"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/dto"
)

type UserUsecase struct {
	repository domain.UserRepository
}

func (u *UserUsecase) Login(request dto.UserLoginRequest) error {
	return nil
}
