package interfaces

import "github.com/auth/service/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) (int, error)
	Login(user domain.User) (domain.User, error)
}
