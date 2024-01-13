package requests

type CreateQuestionRequest struct {
	Name     string `json:"Name" binding:"required"`
	Surname  string `json:"Surname" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Question string `gorm:"column:question" json:"Question"`
}
