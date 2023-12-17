package models

import "time"

type Admin struct {
	ID        int       `gorm:"primarykey"`
	Name      string    `gorm:"column:name" json:"Name"`
	Email     string    `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null;column:email" json:"Email"`
	Password  string    `gorm:"column:password" json:"Password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt"`
}

func (Admin) TableName() string {
	return "admin"
}
