package services

import (
	"klijentske-tehnologije/models"
	"klijentske-tehnologije/repositories"
	"klijentske-tehnologije/requests"
)

type QuestionService interface {
	Create(question requests.CreateQuestionRequest)
}

type QuestionServiceImpl struct {
	QuestionRepository repositories.QuestionRepository
}

// NewQuestionServiceImpl creates a new instance of QuestionServiceImpl.
func NewQuestionServiceImpl(questionRepository repositories.QuestionRepository) QuestionService {
	return &QuestionServiceImpl{
		QuestionRepository: questionRepository,
	}
}

// Create creates a new question based on the provided CreateQuestionRequest.
func (q *QuestionServiceImpl) Create(question requests.CreateQuestionRequest) {
	// Create a QuestionModel based on the CreateQuestionRequest
	questionModel := models.Question{
		Name:     question.Name,
		Surname:  question.Surname,
		Email:    question.Email,
		Question: question.Question,
	}
	// Call the QuestionRepository's Create method to persist the question in the database
	q.QuestionRepository.Create(questionModel)
}
