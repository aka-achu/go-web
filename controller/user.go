package controller

import (
	"github.com/aka-achu/go-web/models"
	"net/http"
)

// UserController is an empty struct on which the handle functions will be implemented
type UserController struct{}

// NewUserController, returns an initialized UserController
func NewUserController() *UserController {
	return &UserController{}
}

// Create returns a handle function to process a user creation request
func (c *UserController) Create(userRepo models.UserRepo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

// Fetch returns a handle function to process a user fetch request
func (c *UserController) Fetch(userRepo models.UserRepo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
