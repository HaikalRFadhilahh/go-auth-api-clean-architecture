package domain

import (
	"time"
)

type User struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Age       int        `json:"age,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type UserRepository interface {
	GetUser(search string, activePage int) ([]User, error)
	GetUserPagination(search string) (int, error)
	GetUserById(id int) (User, error)
	GetUserByUsername(username string) (User, error)
	CreateUser(*User) error
	UpdateUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
}
