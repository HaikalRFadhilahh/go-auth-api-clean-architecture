package domain

import "time"

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type UserRepository interface {
	GetUser(q string) ([]User, error)
	GetUserById(id int) (User, error)
	GetUserByUsername(username string) (User, error)
	CreateUser(*User) error
	UpdateUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
}
