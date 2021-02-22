package data

import (
	"github.com/lucabecci/questions-golang-API/pkg/question"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	Database *gorm.DB
}

type Question struct {
	Title       string
	Description string
	UserID      uint
}

func (q *QuestionRepository) Create(ques Question) (question.Question, bool) {
	newQuestion := question.Question{
		Title:       ques.Title,
		Description: ques.Description,
		UserID:      ques.UserID,
	}

	result := q.Database.Create(&newQuestion)

	if result.RowsAffected < 1 {
		return question.Question{}, false
	}
	result.Scan(&newQuestion)

	return newQuestion, true
}

func (q *QuestionRepository) GetByUser(userID uint) ([]question.Question, bool) {
	var questions []question.Question
	result := q.Database.Where("user_id = ?", userID).Find(&questions)
	result.Scan(&questions)
	if len(questions) < 1 {
		return []question.Question{}, false
	}
	return questions, true
}
