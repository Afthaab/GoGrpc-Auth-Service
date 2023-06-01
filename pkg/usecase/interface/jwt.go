package interfaces

import (
	"github.com/auth/service/pkg/domain"
	"github.com/golang-jwt/jwt"
)

type JwtUseCase interface {
	GenerateAccessToken(userid int, email string, role string) (string, error)
	VerifyToken(token string) (bool, *domain.JWTClaims)
	GetTokenFromString(signedToken string, claims *domain.JWTClaims) (*jwt.Token, error)
}
