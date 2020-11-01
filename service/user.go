package service

import (
	"context"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/svc_error"
	"github.com/aka-achu/go-web/utility"
	"github.com/google/uuid"
)

type UserService struct{}

// NewUserService returns an UserService object
func NewUserService() *UserService {
	return &UserService{}
}

// Create, registers a new user in the application
func (*UserService) Create(user *models.User, userRepo models.UserRepo, ctx context.Context) (*models.User, error) {

	// Extracting the traceID from the context
	traceID := ctx.Value("trace_id").(string)

	// Checking for the existence of the user with the requested user_name
	if userRepo.Exists(user.UserName, ctx) {
		logging.AppLogger.Warnf("Requested user_name already exists in the application. TraceID-%s", traceID)
		return nil, svc_error.ErrUserAlreadyExists
	}

	// Generating a unique id for the user object
	user.ID = uuid.New().String()
	// Hashing the user password
	user.Password = utility.Hash([]byte(user.Password))

	// Creating a record of the user object
	if err := userRepo.Create(user, ctx); err != nil {
		logging.RepoLogger.Errorf("Failed to create the request user. Error-%v TraceID-%s",
			err, traceID)
		return nil, svc_error.ErrFailedToCreateUser
	} else {
		logging.RepoLogger.Infof("Successfully created the requested user. TraceID-%s", traceID)
		user.Password = ""
		return user, nil
	}
}

// Fetch, retrieves the requested user details
func (*UserService) Fetch(userName string, userRepo models.UserRepo,  ctx context.Context) (*models.User, error) {

	// Extracting the traceID from the context
	traceID := ctx.Value("trace_id").(string)

	// Fetching the requested user details
	if user, err := userRepo.Fetch(userName, ctx); err != nil {
		logging.RepoLogger.Errorf("Failed to fetch the request user details. Error-%v TraceID-%s",
			err, traceID)
		return nil, svc_error.ErrFailedToFetchUserDetails
	} else {
		logging.RepoLogger.Infof("Successfully fetched the requested user. TraceID-%s", traceID)
		user.Password = ""
		return user, nil
	}
}
