package svc_error

import "errors"

var (
	ErrUserDoesNotExists        = errors.New("user does not exist in the application")
	ErrFailedToFetchUserDetails = errors.New("failed to fetch user details")
	ErrInvalidCredential        = errors.New("invalid login credential")
	ErrCreatingAccessToken      = errors.New("failed to create access token")
	ErrUserAlreadyExists        = errors.New("user already exists in the application")
	ErrFailedToCreateUser       = errors.New("failed to create the user")
)
