package models

import (
	"fmt"
	"net/http"
	"os"
)

// User, object structure of the user entity
type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	EmployeeID string `json:"employee_id"`
	Age        int    `json:"age"`
}

// UserRepo is a template for the user repo method implementation
type UserRepo interface {
	// Creation of user given an object of user
	Create(user *User) error
	// Fetching the user object given the id of the user
	Fetch(userID string) (*User, error)
}

// UserController is a template for the user controller method implementation
type UserController interface {
	Create(UserRepo) http.HandlerFunc
	Fetch(UserRepo) http.HandlerFunc
}

// TableName return the fully qualified table name for user object
func (*User) TableName() string {
	return fmt.Sprintf("%s.%s", os.Getenv("DB_SCHEMA"), "web_user")
}