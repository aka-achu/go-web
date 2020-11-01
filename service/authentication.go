package service

import (
	"context"
	"errors"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/svc_error"
	"github.com/aka-achu/go-web/utility"
	"gorm.io/gorm"
)

type AuthenticationService struct{}

// NewAuthenticationService returns an AuthenticationService object
func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

// Login, validates the user credentials, on successful
// validation a JWT token will be created which will be
// used to authenticate the client request
func (*AuthenticationService) Login(
	authRequest *models.AuthenticationRequest,
	userRepo models.UserRepo,
	ctx context.Context,
) (*models.AuthenticationResponse, error) {

	// Extracting the traceID from the context
	traceID := ctx.Value("trace_id").(string)

	// Fetching the requested user details
	user, err := userRepo.Fetch(authRequest.UserName, ctx)

	// Checking for the existence of the requested user in the application
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logging.ServiceLogger.Errorf("Request user does not exist in the application. Error-%v TraceID-%s",
			err, traceID)
		return nil, svc_error.ErrUserDoesNotExists
	}
	if err != nil {
		logging.ServiceLogger.Errorf("Failed to fetch the request user details. Error-%v TraceID-%s",
			err, traceID)
		return nil, svc_error.ErrFailedToFetchUserDetails
	}

	// Validating the password
	if utility.Hash([]byte(authRequest.Password)) != user.Password {
		return nil, svc_error.ErrInvalidCredential
	} else {
		// Generating an access_token for the verified user
		accessToken, err := utility.CreateToken(authRequest.UserName)
		if err != nil {
			logging.ServiceLogger.Errorf("Failed to generate an access_token. Error-%v TraceID-%s", err, traceID)
			return nil, svc_error.ErrCreatingAccessToken
		}
		return &models.AuthenticationResponse{
			AuthenticationStatus: true,
			AccessToken:          accessToken,
		}, nil
	}
}
