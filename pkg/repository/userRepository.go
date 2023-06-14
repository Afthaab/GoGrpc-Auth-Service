package repository

import (
	"fmt"

	"github.com/auth/service/pkg/domain"
	repo "github.com/auth/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func (r *userDatabase) FindByUserName(user domain.User) (domain.User, error) {
	result := r.DB.First(&user, "username LIKE ?", user.Username).Error
	return user, result

}
func (r *userDatabase) FindByUserEmail(user domain.User) (domain.User, error) {
	result := r.DB.First(&user, "email LIKE ?", user.Email).Error
	return user, result
}

func (r *userDatabase) Create(user domain.User) error {
	result := r.DB.Create(&user).Error
	return result
}

func (r *userDatabase) FindUserByOtp(user domain.User) (domain.User, error) {
	result := r.DB.Where("otp LIKE ?", user.Otp).First(&user)
	return user, result.Error
}

func (r *userDatabase) NullTheOtp(user domain.User) int64 {
	var userData domain.User
	result := r.DB.Model(&userData).Where("id = ?", user.Id).Update("otp", nil)
	return result.RowsAffected
}

func (r *userDatabase) FindUserById(userid uint) (domain.User, error) {
	user := domain.User{}
	result := r.DB.First(&user, "id = ?", userid).Error
	return user, result
}

func (r *userDatabase) IsOtpVerified(username string) string {
	var otp string
	r.DB.Raw("select otp from users where username LIKE ?", username).Scan(&otp)
	fmt.Println("===========================", otp)
	return otp
}

func (r *userDatabase) DeleteUser(user domain.User) error {
	// var userr domain.User
	// result := r.DB.Where("email LIKE ?", user.Email).Delete(&userr).Error
	// fmt.Println("///////////////////////////////", result)
	// return result
	result := r.DB.Exec("DELETE FROM users WHERE email LIKE ?", user.Email).Error
	fmt.Println("///////////////////////////////", result)
	return result

}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userDatabase{
		DB: db,
	}
}
