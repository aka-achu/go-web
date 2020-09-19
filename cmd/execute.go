package cmd

import (
	"github.com/aka-achu/go-web/controller"
	"github.com/aka-achu/go-web/middleware"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/repo"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Execute() {

	router := mux.NewRouter()
	router.Use(
		middleware.NoAuthLogging,
	)
	router.Use(
		cors.AllowAll().Handler,
	)

	db, err := repo.GetConnection()
	if err != nil {
		panic(err)
	}

	InitUserRoute(
		router,
		controller.NewUserController(),
		repo.NewUserRepo(db),
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

func InitUserRoute(r *mux.Router, userController models.UserController, userRepo models.UserRepo) {
	var userRouter = r.PathPrefix("/api/v1/user").Subrouter()
	userRouter.HandleFunc("/create", userController.Create(userRepo))
	userRouter.HandleFunc("/fetch", userController.Fetch(userRepo))
}
