package controller

import (
	"encoding/json"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/response"
	"github.com/aka-achu/go-web/utility"
	"github.com/gorilla/mux"
	"net/http"
)

// UserController is an empty struct on which the handle functions will be implemented
type UserController struct{}

// NewUserController, returns an initialized UserController
func NewUserController() *UserController {
	return &UserController{}
}

// Create returns a handle function to process a user creation request
func (c *UserController) Create(userRepo models.UserRepo, userService models.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userCreationRequest models.User
		// Getting the request tracing id from the request context
		requestTraceID := r.Context().Value("trace_id").(string)

		// Decoding the request body data to models.User object
		err := json.NewDecoder(r.Body).Decode(&userCreationRequest)
		if err != nil {
			logging.AppLogger.Errorf("Failed to decode the request body. Error-%v TraceID-%s",
				err, requestTraceID)
			response.BadRequest(w, "100", err.Error())
			return
		}

		// Validating the request object for required fields
		err = utility.Validate.Struct(userCreationRequest)
		if err != nil {
			logging.AppLogger.Errorf("Failed to validate the request body. Error-%v TraceID-%s",
				err, requestTraceID)
			response.BadRequest(w, "101", err.Error())
			return
		}

		// Using the user creation service to create the requested user for the application
		if user, err := userService.Create(&userCreationRequest, userRepo, r.Context()); err != nil {
			logging.AppLogger.Errorf("Failed to create the requested user. Error-%v TraceID-%s",
				err, requestTraceID)
			response.InternalServerError(w, "102", err.Error())
		} else {
			logging.AppLogger.Infof("Successfully created the requested user. TraceID-%s",
				requestTraceID)
			response.Success(w, "103","Successful creation of requested user", user)
		}
	}
}

// Fetch returns a handle function to process a user fetch request
func (c *UserController) Fetch(userRepo models.UserRepo, userService models.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting the request tracing id from the request context
		requestTraceID := r.Context().Value("trace_id").(string)

		// Extracting the request user_name from the request URI
		userName := mux.Vars(r)["user_name"]

		// Fetching the requested user details
		if user, err := userService.Fetch(userName, userRepo,  r.Context()); err != nil {
			logging.AppLogger.Errorf("Failed to fetch the requested user. Error-%v TraceID-%s",
				err, requestTraceID)
			response.InternalServerError(w, "104", err.Error())
		} else {
			logging.AppLogger.Infof("Successfully fetched the requested user. TraceID-%s",
				requestTraceID)
			response.Success(w, "103","Successful fetch of requested user data", user)
		}


	}
}
