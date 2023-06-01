package interfaces

import "github.com/auth/service/pkg/domain"

type UserRepo interface {
	FindUser(user domain.User) (int, error)
	Create(user domain.User) (int, error)
	FindByUserName(name string) (domain.User, error)
	FindByUserEmail(email string) (domain.User, error)
}
