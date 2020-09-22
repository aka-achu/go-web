package utility

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Initialize() {
	Validate = validator.New()
}