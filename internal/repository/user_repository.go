package repository

import (
	"database/sql"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUser(q string) ([]domain.User, error) {
	return nil, nil
}

func (u *userRepository) GetUserById(id int) (domain.User, error) {
	return domain.User{}, nil
}

func (u *userRepository) GetUserByUsername(username string) (domain.User, error) {
	return domain.User{}, nil
}

func (u *userRepository) CreateUser(domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (u *userRepository) UpdateUser(user domain.User, id int) (domain.User, error) {
	return domain.User{}, nil
}

func (u *userRepository) DeleteUser(id int) (domain.User, error) {
	return domain.User{}, nil
}
