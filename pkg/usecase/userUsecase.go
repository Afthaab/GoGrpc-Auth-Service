package usecase

import (
	"errors"

	domain "github.com/auth/service/pkg/domain"
	repo "github.com/auth/service/pkg/repository/interface"
	useCase "github.com/auth/service/pkg/usecase/interface"
	utility "github.com/auth/service/pkg/utils"
)

type userUseCase struct {
	Repo repo.UserRepo
}

func (u *userUseCase) Register(user domain.User) (int, error) {
	var id int
	var err error
	id, err = u.Repo.FindUser(user)
	if err != nil {
		return id, err
	}
	user.Password = utility.HashPassword(user.Password)
	id, err = u.Repo.Create(user)
	if err != nil {
		return id, err
	}
	return id, nil

}
func (u *userUseCase) Login(user domain.User) (domain.User, error) {
	if user.Username != "" {
		userDetails, err := u.Repo.FindByUserName(user.Username)
		if err != nil {
			return userDetails, errors.New("Could not find the user")
		}
		if !utility.VerifyPassword(user.Password, userDetails.Password) {
			return userDetails, errors.New("Password in worng or did not match")
		}
		return userDetails, nil
	} else if user.Email != "" {
		userDetails, err := u.Repo.FindByUserEmail(user.Email)
		if err != nil {
			return userDetails, errors.New("Could not find the user")
		}
		if !utility.VerifyPassword(user.Password, userDetails.Password) {
			return userDetails, errors.New("Password in worng or did not match")
		}
		return userDetails, nil
	}
	return user, nil
}

func NewUserUseCase(repo repo.UserRepo) useCase.UserUseCase {
	return &userUseCase{
		Repo: repo,
	}
}
