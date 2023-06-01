package repository

import (
	"errors"

	"github.com/auth/service/pkg/domain"
	repo "github.com/auth/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func (r *userDatabase) FindUser(user domain.User) (int, error) {
	dbErr := r.DB.Raw("select * from users where username LIKE ?", user.Username).Error
	// dbErr := r.DB.Where("where username = ?", user.Username).First(&user).Error
	if dbErr != nil {
		return 0, errors.New("Username Already exists")
	}
	dbErr = r.DB.Raw("select * from users where email LIKE ?", user.Email).Error
	// dbErr = r.DB.Where("where email = ?", user.Email).First(&user).Error
	if dbErr != nil {
		return 0, errors.New("email Already exists")
	}
	if dbErr == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return int(user.Id), dbErr

}

func (r *userDatabase) Create(user domain.User) (int, error) {
	result := r.DB.Create(&user)
	return int(user.Id), result.Error
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
