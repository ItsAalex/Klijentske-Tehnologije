package responses

type UserResponse struct {
	ID           int    `json:"ID"`
	Firstname    string `gorm:"not null" json:"Firstname" binding:"required,min=2,max=50"`
	Lastname     string `gorm:"not null" json:"Lastname" binding:"required,min=2,max=50"`
	Email        string `gorm:"not null"`
	StudyProgram string `gorm:"not null"`
	Index        string `gorm:"not null"`
	Comment      string `gorm:"not null"`
}
