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

func (u *userUseCase) Register(user domain.User) error {
	// Validating the JSON using Validator Pacakge
	validationErr := utility.ValidateUser(user)
	if validationErr != nil {
		return validationErr
	}

	// Searching in Database for Existing User Credentials
	_, err := u.Repo.FindByUserEmail(user)
	if err == nil {
		return errors.New("Email Address already exists")
	}
	_, err = u.Repo.FindByUserName(user)
	if err == nil {
		return errors.New("Username Already exists")
	}

	// Generating OTP for the User
	otp := utility.Otpgeneration(user.Email)
	user.Otp = otp

	// Hashing the Password
	user.Password = utility.HashPassword(user.Password)

	// Creating the user
	err = u.Repo.Create(user)
	if err != nil {
		return err
	}
	return nil

}

func (u *userUseCase) RegisterValidate(user domain.User) error {
	// searching for the user with otp
	user, err := u.Repo.FindUserByOtp(user)
	if err != nil {
		return errors.New("Entered wrong OTP")
	}

	// Null the otp field
	rows := u.Repo.NullTheOtp(user)
	if rows == 0 {
		return errors.New("Could not update the OTP")
	}
	return nil

}

func (u *userUseCase) Login(user domain.User) (domain.User, error) {
	if user.Username != "" { // check the user in the database through username
		userDetails, err := u.Repo.FindByUserName(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}
		if !utility.VerifyPassword(user.Password, userDetails.Password) {
			return userDetails, errors.New("Password in worng or did not match")
		}
		return userDetails, nil
	} else if user.Email != "" { // check the user in the database through email
		userDetails, err := u.Repo.FindByUserEmail(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}
		if !utility.VerifyPassword(user.Password, userDetails.Password) {
			return userDetails, errors.New("Password in worng or did not match")
		}
		return userDetails, nil
	}
	return user, nil
}

func (u *userUseCase) FindUserByID(userid uint) (domain.User, error) {
	user, err := u.Repo.FindUserById(userid)
	if err != nil {
		return user, errors.New("Unauthorized User")
	}
	return user, nil
}

func NewUserUseCase(repo repo.UserRepo) useCase.UserUseCase {
	return &userUseCase{
		Repo: repo,
	}
}
