package user

import (
	helpers "klijentske-tehnologije/helpers"
	requests "klijentske-tehnologije/requests"
	responses "klijentske-tehnologije/responses"
	services "klijentske-tehnologije/services"
	http "net/http"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

type QuestionController struct {
	questionService services.QuestionService
	Db              *gorm.DB
}

func NewQuestionController(service services.QuestionService) *QuestionController {
	return &QuestionController{
		questionService: service,
	}
}

func (ctrl *QuestionController) Create(c *gin.Context) {
	// Logic for creating a question in the database
	createQuestionRequest := requests.CreateQuestionRequest{}
	err := c.ShouldBindJSON(&createQuestionRequest)
	helpers.ErrorPanic(err)

	ctrl.questionService.Create(createQuestionRequest)
	webResponse := responses.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   nil,
	}
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}
