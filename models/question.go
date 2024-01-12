package models

import "time"

type Question struct {
	ID        int       `gorm:"primarykey"`
	Name      string    `gorm:"column:name" json:"Name"`
	Surname   string    `gorm:"column:surname" json:"Surname"`
	Email     string    `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null;column:email" json:"Email"` //unique, do I need all of this?
	Question  string    `gorm:"column:question" json:"Question"`
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt"`
}

func (Question) TableName() string {
	return "question"
}
