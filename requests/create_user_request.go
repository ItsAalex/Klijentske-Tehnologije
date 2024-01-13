package requests

type CreateUserRequest struct {
	Firstname    string `json:"FirstName" binding:"required,min=2,max=50"`
	Lastname     string `json:"LastName" binding:"required,min=2,max=50"`
	Email        string `json:"Email" binding:"required,email"`
	StudyProgram string
	Index        string `json:"Index" binding:"required"`
	Comment      string `json:"Comment" binding:"required"`
}
