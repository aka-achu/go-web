package service

import (
	"errors"
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"github.com/aka-achu/go-web/utility"
	"github.com/google/uuid"
)

type UserService struct{}

// NewUserService returns an UserService object
func NewUserService() *UserService {
	return &UserService{}
}

// Create, registers a new user in the application
func (*UserService) Create(user *models.User, userRepo models.UserRepo, traceID string) (*models.User, error) {

	// Checking for the existence of the user with the requested user_name
	if userRepo.Exists(user.UserName) {
		logging.AppLogger.Warnf("Requested user_name already exists in the application. TraceID-%s", traceID)
		return nil, errors.New("user_name already exists in the application")
	}

	// Generating a unique id for the user object
	user.ID = uuid.New().String()
	// Hashing the user password
	user.Password = utility.Hash([]byte(user.Password))

	// Creating a record of the user object
	if err := userRepo.Create(user); err != nil {
		logging.RepoLogger.Errorf("Failed to create the request user. Error-%v TraceID-%s",
			err, traceID)
		return nil, err
	} else {
		logging.RepoLogger.Infof("Successfully created the requested user.TraceID-%s", traceID)
		user.Password = ""
		return user, nil
	}
}
