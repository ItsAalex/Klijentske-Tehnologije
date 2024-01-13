package models

import (
	time "time"
)

type User struct {
	ID           int       `gorm:"primarykey"`
	Firstname    string    `gorm:"column:first_name" json:"Firstname"`
	Lastname     string    `gorm:"column:last_name" json:"Lastname"`
	Email        string    `gorm:"column:email;unique" json:"Email"`
	StudyProgram string    `gorm:"column:study_program" json:"StudyProgram"`
	Index        string    `gorm:"column:index" json:"Index"`
	Comment      string    `gorm:"column:comment" json:"Comment" `
	CreatedAt    time.Time `gorm:"column:created_at" json:"CreatedAt"`
}

func (User) TableName() string {
	return "user"
}
