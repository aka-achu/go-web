package utility

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Initialize() {
	Validate = validator.New()
}

func Hash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
