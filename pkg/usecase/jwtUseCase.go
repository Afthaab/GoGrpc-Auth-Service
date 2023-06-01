package usecase

import (
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
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(500)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(u.SecretKey))

	return accessToken, err

}

func NewJWTUseCase() interfaces.JwtUseCase {
	return &jwtUseCase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
