package repositories

import (
	"klijentske-tehnologije/helpers"
	"klijentske-tehnologije/models"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(question models.Question) *models.Question
}
type QuestionRepositoryImpl struct {
	Db *gorm.DB
}

func NewQuestionRepositoryImpl(Db *gorm.DB) QuestionRepository {
	return &QuestionRepositoryImpl{Db: Db}
}

func (q *QuestionRepositoryImpl) Create(question models.Question) *models.Question {
	result := q.Db.Save(&question)
	helpers.ErrorPanic(result.Error)

	return &question
}
