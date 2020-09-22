package utility

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"
)


// Hash returns a sha1 hash value of the requested data
func Hash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// CreateToken, generates a JWT token
func CreateToken(userName string) (string, error) {
	// Adding data ot the jwt claims
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = userName

	// The validity of the JWT token will be one hour
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generating JWT using the claims, secret token present in the env variables and the signing method
	return jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims).SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

// VerifyToken, validates integrity of a JWT token
func VerifyToken(bearToken string) error {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		return errors.New("invalid bearToken format")
	}
	if token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	}); err != nil {
		return err
	} else {
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			return errors.New("invalid access token")
		} else {
			return nil
		}
	}
}
