package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(db *gorm.DB) ([]User, error)
	GetOne(db *gorm.DB, id uint) (User, error)
	GetByEmail(db *gorm.DB, email string) (User, error)
	Create(db *gorm.DB, user *User) error
	Delete(db *gorm.DB, id uint) error
	Update(db *gorm.DB, id uint, user User) (User, error)
}
