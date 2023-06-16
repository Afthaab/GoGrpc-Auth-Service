package interfaces

import "github.com/auth/service/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) error
	RegisterValidate(user domain.User) error
	Login(user domain.User) (domain.User, error)
	ValidateJwtUser(userid uint) (domain.User, error)
	AdminLogin(user domain.User) (domain.User, error)
	ForgotPassword(user domain.User) error
}
