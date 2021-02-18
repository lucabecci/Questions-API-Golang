package question

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	UserID      uint
}
