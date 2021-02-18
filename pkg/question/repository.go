package question

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(db *gorm.DB, userID uint) ([]Question, error)
	GetOne(db *gorm.DB, id uint) (Question, error)
	Create(db *gorm.DB, userID uint) error
	Delete(db *gorm.DB, id uint) error
	Update(db *gorm.DB, id uint, question Question) (Question, error)
}
