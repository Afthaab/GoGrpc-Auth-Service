package domain

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	Userid uint
	Email  string
	Source string
	jwt.StandardClaims
}
