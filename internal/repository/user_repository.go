package repository

import (
	"database/sql"
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
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
	// Building Query
	query := "SELECT id,name,age,username FROM users where id=?"

	// Create Variabel To Temp Data User Models Query
	var data domain.User

	// Exec Query Row
	err := u.db.QueryRow(query, id).Scan(&data.ID, &data.Name, &data.Age, &data.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, apierror.ErrForbidden
		}

		return data, apierror.APIErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	return data, nil
}

func (u *userRepository) GetUserByUsername(username string) (domain.User, error) {
	// Build Query
	query := "SELECT id,name,username,password from users WHERE username = ?"

	// Create Var Data Handler
	var data domain.User

	// Exec Query
	err := u.db.QueryRow(query, username).Scan(&data.ID, &data.Name, &data.Username, &data.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, apierror.APIErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Status:     "error",
				Message:    "Username or Password Wrong",
			}
		}
		return domain.User{}, apierror.APIErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	return data, nil
}

func (u *userRepository) CreateUser(data *domain.User) error {
	// Build Insert Query
	query := "INSERT INTO users(name,age,username,password) VALUES (?,?,?,?)"

	// Exec Query
	_, err := u.db.Exec(query, data.Name, data.Age, data.Username, data.Password)
	if err != nil {
		return apierror.APIErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	// Return
	return nil
}

func (u *userRepository) UpdateUser(user domain.User, id int) (domain.User, error) {
	return domain.User{}, nil
}

func (u *userRepository) DeleteUser(id int) (domain.User, error) {
	return domain.User{}, nil
}
