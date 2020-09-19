package controller

import (
	"github.com/aka-achu/go-web/models"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Create(userRepo models.UserRepo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (c *UserController) Fetch(userRepo models.UserRepo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
