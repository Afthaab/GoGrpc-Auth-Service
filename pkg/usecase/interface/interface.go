package interfaces

import "github.com/auth/service/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) error
	RegisterValidate(user domain.User) error
	Login(user domain.User) (domain.User, error)
	FindUserByID(userid uint) (domain.User, error)
}
