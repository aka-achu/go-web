package controller

import (
	"encoding/json"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/response"
	"github.com/aka-achu/go-web/utility"
	"net/http"
)

// AuthController is an empty struct on which the handle functions will be implemented
type AuthController struct{}

// NewAuthController, returns an initialized AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// Create returns a handle function to process a user creation request
func (c *UserController) Login(userRepo models.UserRepo, authService models.AuthenticationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest models.AuthenticationRequest
		// Getting the request tracing id from the request context
		requestTraceID := r.Context().Value("trace_id").(string)

		// Decoding the request body data to models.User object
		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			logging.AppLogger.Errorf("Failed to decode the request body. Error-%v TraceID-%s",
				err, requestTraceID)
			response.BadRequest(w, "100", err.Error())
			return
		}

		// Validating the request object for required fields
		if err := utility.Validate.Struct(loginRequest); err != nil {
			logging.AppLogger.Errorf("Failed to validate the request body. Error-%v TraceID-%s",
				err, requestTraceID)
			response.BadRequest(w, "101", err.Error())
			return
		}

		if loginResponse, err := authService.Login(&loginRequest, userRepo, r.Context()); err != nil {
			logging.AppLogger.Errorf("User login failed. Error-%v TraceID-%s",
				err, requestTraceID)
			response.InternalServerError(w, "201", err.Error())
		} else {
			logging.AppLogger.Errorf("User login successful. TraceID-%s",
				requestTraceID)
			response.Success(w, "202","Login successful", loginResponse)
		}

	}
}