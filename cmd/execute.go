package cmd

import (
	"github.com/aka-achu/go-web/controller"
	"github.com/aka-achu/go-web/middleware"
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

	var sampleController *controller.Sample
	var sampleRouter = router.PathPrefix("/api/v1/sample").Subrouter()

	sampleRouter.HandleFunc("/hello", sampleController.HelloWorld)

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
