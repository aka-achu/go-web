package cmd

import (
	"context"
	"github.com/aka-achu/go-web/controller"
	"github.com/aka-achu/go-web/repo"
	"github.com/aka-achu/go-web/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	controller.InitUserRoute(
		router,
		controller.NewUserController(),
		repo.NewUserRepo(db),
		service.NewUserService(),
	)
	controller.InitAuthenticationRoute(
		router,
		controller.NewUserController(),
		repo.NewUserRepo(db),
		service.NewAuthenticationService(),
	)
	server := &http.Server{
		Addr:              os.Getenv("SERVER_ADDRESS"),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Println("Server has started at", os.Getenv("SERVER_ADDRESS"))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", os.Getenv("SERVER_ADDRESS"), err)
	}
	<-done
	log.Println("Server has stopped")

}


