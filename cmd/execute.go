package cmd

import (
	"github.com/aka-achu/go-web/controller"
	"github.com/aka-achu/go-web/middleware"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/repo"
	"github.com/aka-achu/go-web/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Execute, initializes the web application
// It creates a router,
//    configures the middlewares,
//    initializes the controllers,
//    initialized the repo layers,
//    initialized the service layers,
//    registers the handle functions,
//    starts the web server.
func Execute() {

	// Creating a new router
	router := mux.NewRouter()
	// Initializing the middlewares
	router.Use(
		middleware.NoAuthLogging,
	)
	// Handling cors access
	router.Use(
		cors.AllowAll().Handler,
	)
	// Getting a new database connection
	db, err := repo.GetConnection()
	if err != nil {
		panic(err)
	}

	// Registering handle functions
	InitUserRoute(
		router,
		controller.NewUserController(),
		repo.NewUserRepo(db),
		service.NewUserService(),
	)
	InitAuthenticationRoute(
		router,
		controller.NewUserController(),
		repo.NewUserRepo(db),
		service.NewAuthenticationService(),
	)

	if os.Getenv("BUILD") == "Prod" {
		log.Fatal(http.ListenAndServeTLS(
			os.Getenv("SERVER_ADDRESS"),
			filepath.Join(os.Getenv("PATH_TO_CERTIFICATE")),
			filepath.Join(os.Getenv("PATH_TO_KEY")),
			router,
		))
	} else {
		log.Fatal(http.ListenAndServe(
			os.Getenv("SERVER_ADDRESS"),
			router,
		))
	}
}

// InitUserRoute, registers user handle function in the given router
func InitUserRoute(
	r *mux.Router,
	userController models.UserController,
	userRepo models.UserRepo,
	userService models.UserService,
) {

	// Creating a sub-router for common path
	var userRouter = r.PathPrefix("/api/v1/user").Subrouter()
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
	authRouter.HandleFunc("/login", authController.Login(userRepo, authService)).Methods("POST")
}
