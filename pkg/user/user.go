package user

import (
	"github.com/lucabecci/questions-golang-API/pkg/question"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User is the main struct of the table
type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Email     string
	Password  string
	Questions []question.Question `gorm:"foreignKey:UserID"`
}

func (u *User) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)

	return nil
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
