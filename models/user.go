package models

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
)

// User, object structure of the user entity
type User struct {
	ID         string `json:"id"`
	UserName   string `json:"user_name"    validate:"required"`
	Password   string `json:"password"     validate:"required"`
	FirstName  string `json:"first_name"   validate:"required"`
	LastName   string `json:"last_name"    validate:"required"`
	EmployeeID string `json:"employee_id"  validate:"required"`
	Age        int    `json:"age"          validate:"required"`
	gorm.Model
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
