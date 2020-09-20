package controller

import (
	"encoding/json"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/response"
	"github.com/aka-achu/go-web/utility"
	"github.com/google/uuid"
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

		userCreationRequest.ID = uuid.New().String()
		userCreationRequest.Password = utility.Hash([]byte(userCreationRequest.Password))

		if err := userRepo.Create(&userCreationRequest); err != nil {
			logging.RepoLogger.Errorf("Failed to create the request user. Error-%v TraceID-%s",
				err, requestTraceID)
			response.BadRequest(w, "102", err.Error())
		} else {
			logging.RepoLogger.Errorf("Successfully created the requested user.TraceID-%s",
				requestTraceID)
			userCreationRequest.Password = ""
			response.Success(w, "103", "Successful creation of user", userCreationRequest)
		}
	}
}

// Fetch returns a handle function to process a user fetch request
func (c *UserController) Fetch(userRepo models.UserRepo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
