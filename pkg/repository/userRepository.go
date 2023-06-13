package repository

import (
	"errors"
	"fmt"

	"github.com/auth/service/pkg/domain"
	repo "github.com/auth/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func (r *userDatabase) FindUser(user domain.User) (int, error) {
	result := r.DB.First(&user, "username LIKE ?", user.Username).RowsAffected
	if result != 0 {
		return 0, errors.New("Username Already Exists !")
	}
	result = r.DB.First(&user, "email LIKE ?", user.Email).RowsAffected
	if result != 0 {
		return 0, errors.New("Email Already Exists !")
	}
	return int(user.Id), nil
}

func (r *userDatabase) Create(user domain.User) (int, error) {
	result := r.DB.Create(&user)
	return int(user.Id), result.Error
}

func (r *userDatabase) FindUserByOtp(user domain.User) (domain.User, error) {
	result := r.DB.Where("otp LIKE ?", user.Otp).First(&user)
	if result.Error != nil {
		return user, errors.New("Wrong OTP Entered")
	}
	return user, nil
}

func (r *userDatabase) NullTheOtp(user domain.User) (int, error) {
	var userData domain.User
	result := r.DB.Model(&userData).Where("id = ?", user.Id).Update("otp", nil)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected != 0 {
		return int(user.Id), nil
	}
	return int(userData.Id), errors.New("Could not delete the otp")
}

func (r *userDatabase) FindByUserName(name string) (domain.User, error) {
	var userData domain.User
	result := r.DB.Raw("select * from users where username = ?", name).Scan(&userData).Error
	return userData, result
}
func (r *userDatabase) FindByUserEmail(email string) (domain.User, error) {
	var userData domain.User
	result := r.DB.Raw("select * from users where email = ?", email).Scan(&userData).Error
	return userData, result
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userDatabase{
		DB: db,
	}
}
