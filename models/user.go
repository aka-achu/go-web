package models

import (
	"context"
	"gorm.io/gorm"
	"net/http"
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
	Create(*User, context.Context) error
	// Fetching the user object given the id of the user
	Fetch(string, context.Context) (*User, error)
	// Checking the existence of the requested user
	Exists(string, context.Context) bool
}

// UserController is a template for the user controller method implementation
type UserController interface {
	Create(UserRepo, UserService) http.HandlerFunc
	Fetch(UserRepo, UserService) http.HandlerFunc
}

// UserService is a template for the user service method implementation
type UserService interface {
	Create(*User, UserRepo, context.Context) (*User, error)
	Fetch(string, UserRepo, context.Context) (*User, error)
}

// TableName return the fully qualified table name for user object
func (*User) TableName() string {
	return "web_user"
}
