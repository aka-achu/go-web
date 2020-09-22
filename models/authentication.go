package models

import (
	"context"
	"net/http"
)

// AuthenticationRequest, object structure contains required fields for user login request
type AuthenticationRequest struct {
	UserName string `json:"user_name"    validate:"required"`
	Password string `json:"password"     validate:"required"`
}

type AuthenticationResponse struct {
	AuthenticationStatus bool   `json:"authentication_status"`
	AccessToken          string `json:"access_token"`
}

// AuthenticationController is a template for the authentication controller method implementation
type AuthenticationController interface {
	Login(UserRepo, AuthenticationService) http.HandlerFunc
}

// AuthenticationService is a template for the authentication service method implementation
type AuthenticationService interface {
	Login(*AuthenticationRequest, UserRepo, context.Context) (*AuthenticationResponse, error)
}
