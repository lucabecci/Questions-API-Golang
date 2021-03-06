package data

import (
	"errors"

	"github.com/lucabecci/questions-golang-API/pkg/question"
	"github.com/lucabecci/questions-golang-API/pkg/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func (u *UserRepository) GetOne(id uint) (user.User, bool) {
	var userReturned user.User

	rows := u.Database.Find(&userReturned, id)

	rows.Scan(&userReturned)

	if userReturned.ID == 0 {
		return user.User{}, false
	}

	return userReturned, true
}

func (u *UserRepository) Create(email string, password string) (user.User, bool) {

	newUser := user.User{
		Email:     email,
		Password:  password,
		Questions: []question.Question{},
	}
	newUser.HashPassword()
	result := u.Database.Create(&newUser)

	if result.RowsAffected < 1 {
		return user.User{}, false
	}
	result.Scan(&newUser)

	return newUser, true
}

func (u *UserRepository) UserExists(email string) (user.User, error) {
	var usr user.User
	result := u.Database.Where("email = ?", email).First(&usr)

	result.Scan(&usr)
	if usr.ID == 0 {
		return user.User{}, errors.New("User Not Found")
	}
	return usr, nil

}

func (u *UserRepository) EmailInUse(email string) bool {
	var usr user.User
	result := u.Database.Where("email = ?", email).Find(&usr)
	if result.RowsAffected == 1 {
		return true
	}
	return false
}
