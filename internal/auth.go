package internal

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

// JWTKey is used to create the signature
var JWTKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims is a struct that will be encoded in a jwt
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
