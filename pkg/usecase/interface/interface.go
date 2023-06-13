package interfaces

import "github.com/auth/service/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) (int, error)
	RegisterValidate(user domain.User) (int, error)
	Login(user domain.User) (domain.User, error)
	FindByUserEmail(email string) (domain.User, error)
}
