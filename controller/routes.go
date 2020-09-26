package controller

import (
	"github.com/aka-achu/go-web/middleware"
	"github.com/aka-achu/go-web/models"
	"github.com/gorilla/mux"
)

// InitUserRoute, registers user handle function in the given router
func InitUserRoute(
	r *mux.Router,
	userController models.UserController,
	userRepo models.UserRepo,
	userService models.UserService,
) {

	// Creating a sub-router for common path
	var userRouter = r.PathPrefix("/api/v1/user").Subrouter()
	userRouter.Use(middleware.NoAuthLogging)
	userRouter.HandleFunc("/create", userController.Create(userRepo, userService)).
		Methods("POST")
	userRouter.HandleFunc("/fetch/{user_name}", userController.Fetch(userRepo, userService)).
		Methods("GET")
}

// InitAuthenticationRoute, registers authentication handle function in the given router
func InitAuthenticationRoute(
	r *mux.Router,
	authController models.AuthenticationController,
	userRepo models.UserRepo,
	authService models.AuthenticationService,
) {

	// Creating a sub-router for common path
	var authRouter = r.PathPrefix("/api/v1/auth").Subrouter()
	authRouter.Use(middleware.AuthLogging)
	authRouter.HandleFunc("/login", authController.Login(userRepo, authService)).Methods("POST")
}