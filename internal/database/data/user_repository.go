package data

import (
	"github.com/lucabecci/questions-golang-API/internal/database"
	"github.com/lucabecci/questions-golang-API/pkg/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

type UserInput struct {
	Email    string
	Password string
}

func GetInstance() *UserRepository {
	userRepository := UserRepository{Database: database.Database()}
	return &userRepository
}

func (u *UserRepository) GetAll() ([]user.User, bool) {
	var users []user.User

	rows := u.Database.Find(&users)

	rows.Scan(&users)

	if len(users) < 1 {
		return []user.User{}, false
	}
	return users, true
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

func (u *UserRepository) Create(newUser *UserInput) (user.User, bool) {

	usr := user.User{
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	result := u.Database.Create(&usr)

	if result.RowsAffected < 1 {
		return user.User{}, false
	}
	return usr, true
}
