package usecase

import (
	"errors"
	"fmt"

	domain "github.com/auth/service/pkg/domain"
	repo "github.com/auth/service/pkg/repository/interface"
	useCase "github.com/auth/service/pkg/usecase/interface"
	utility "github.com/auth/service/pkg/utils"
)

type userUseCase struct {
	Repo repo.UserRepo
}

// -------------------------- User Authentication -----------------------------

func (u *userUseCase) Register(user domain.User) error {
	// Validating the JSON using Validator Pacakge
	validationErr := utility.ValidateUser(user)
	if validationErr != nil {
		return validationErr
	}

	// Searching in Database for Existing User Credentials
	userData, err := u.Repo.FindByUserEmail(user)

	if err == nil {
		// Deleting the User if he already exits but not verfied
		if userData.Isverified == false {
			errs := u.Repo.DeleteUser(user)
			if errs != nil {
				return errors.New("Could not delete unethenticated user")
			}
		} else {
			return errors.New("Email Address already exists")
		}
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

	// inserting avatar
	user.Profile = "https://images-for-deliveryapp.s3.ap-south-1.amazonaws.com/5.jpg"

	// Creating the user
	err = u.Repo.Create(user)
	if err != nil {
		return err
	}
	return nil

}

func (u *userUseCase) RegisterValidate(user domain.User) (domain.User, error) {
	// searching for the user with otp
	user, err := u.Repo.FindUserByOtp(user)
	if err != nil {
		return user, errors.New("Entered wrong OTP")
	}

	// Null the otp field
	rows := u.Repo.NullTheOtp(user)
	if rows == 0 {
		return user, errors.New("Could not update the OTP")
	}

	// User gets Verified
	user, err = u.Repo.VerifyUser(user)
	if err != nil {
		return user, errors.New("Could not verifiy the user")
	}
	return user, nil

}

func (u *userUseCase) Login(user domain.User) (domain.User, error) {
	var userDetails domain.User
	var err error
	if user.Username != "" { // check the user in the database through username

		userDetails, err = u.Repo.FindByUserName(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}

	} else if user.Email != "" { // check the user in the database through email

		userDetails, err = u.Repo.FindByUserEmail(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}
	}

	// Deleting the User if he already exits but not verfied
	fmt.Println(userDetails)
	if userDetails.Isverified == false {
		errs := u.Repo.DeleteUser(userDetails)
		if errs != nil {
			return userDetails, errors.New("Could not delete unethenticated user")
		}
		return userDetails, errors.New("User not Authenticated, Register again")
	}

	// checking the hashed password
	if !utility.VerifyPassword(user.Password, userDetails.Password) {
		return userDetails, errors.New("Password in worng or did not match")
	}

	return userDetails, nil
}

func (u *userUseCase) ForgotPassword(user domain.User) error {
	user, err := u.Repo.FindByUserEmail(user)
	if err != nil {
		return errors.New("Email Address not found!")
	}

	// Generating OTP for the User
	otp := utility.Otpgeneration(user.Email)
	user.Otp = otp

	err = u.Repo.UpdateOtp(user)
	if err != nil {
		return errors.New("Could not update the OTP")
	}
	return nil
}

func (u *userUseCase) ChangePassword(user domain.User) error {
	// Hashing the Password
	user.Password = utility.HashPassword(user.Password)
	err := u.Repo.ChangePassword(user)
	if err != nil {
		return errors.New("Could not change the password")
	}
	return nil
}

// -------------------------- Jwt User Authentication -----------------------------

func (u *userUseCase) ValidateJwtUser(userid uint) (domain.User, error) {
	user, err := u.Repo.FindUserById(userid)
	if err != nil {
		return user, errors.New("Unauthorized User")
	}
	return user, nil
}

// -------------------------- Admin Authentication -----------------------------

func (u *userUseCase) AdminLogin(user domain.User) (domain.User, error) {
	if user.Username != "" {

		// check the user in the database through username
		userDetails, err := u.Repo.FindByUserName(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}

		// checking if the user has admin access
		if userDetails.Isadmin == false {
			return userDetails, errors.New("Admin access not found")
		}

		// checking the hashed password
		if !utility.VerifyPassword(user.Password, userDetails.Password) {
			return userDetails, errors.New("Password in worng or did not match")
		}
		return userDetails, nil

	} else if user.Email != "" {

		// check the user in the database through email
		userDetails, err := u.Repo.FindByUserEmail(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}

		// checking the hashed password
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
