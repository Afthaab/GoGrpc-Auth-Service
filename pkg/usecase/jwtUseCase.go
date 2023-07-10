package usecase

import (
	"fmt"
	"os"
	"time"

	domain "github.com/auth/service/pkg/domain"
	interfaces "github.com/auth/service/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
)

type jwtUseCase struct {
	SecretKey string
}

func (u *jwtUseCase) GenerateAccessToken(userid int, email string, role string) (string, error) {
	claims := domain.JWTClaims{
		Userid: uint(userid),
		Email:  email,
		Source: "AccessToken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().AddDate(1, 0, 0).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(u.SecretKey))

	return accessToken, err

}

func (u *jwtUseCase) VerifyToken(token string) (bool, *domain.JWTClaims) {
	claims := &domain.JWTClaims{}
	tkn, err := u.GetTokenFromString(token, claims)
	if err != nil {
		return false, claims
	}
	if tkn.Valid {
		if err := claims.Valid(); err != nil {
			return false, claims
		}
	}
	return true, claims

}

func (u *jwtUseCase) GetTokenFromString(signedToken string, claims *domain.JWTClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(u.SecretKey), nil
	})
}

func NewJWTUseCase() interfaces.JwtUseCase {
	return &jwtUseCase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
