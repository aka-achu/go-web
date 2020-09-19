package models

import (
	"net/http"
)

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	EmployeeID string `json:"employee_id"`
	Age        int    `json:"age"`
}

type UserRepo interface {
	Create(user *User) error
	Fetch(userID string) (*User, error)
}

type UserController interface {
	Create(UserRepo) http.HandlerFunc
	Fetch(UserRepo) http.HandlerFunc
}
