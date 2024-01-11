package models

import (
	time "time"
)

type User struct {
	ID        int       `gorm:"primarykey"`
	Email     string    `gorm:"column:email;unique" json:"Email"`
	Password  string    `gorm:"column:password;type:varchar(255)" json:"Password"`
	Firstname string    `gorm:"column:first_name" json:"Firstname"`
	Lastname  string    `gorm:"column:last_name" json:"Lastname"`
	Address   string    `gorm:"column:address" json:"Address"`
	City      string    `gorm:"column:city" json:"City"`
	Postcode  string    `gorm:"column:post_code" json:"PostCode"`
	Phone     string    `gorm:"column:phone" json:"Phone"`
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt"`
}

func (User) TableName() string {
	return "user"
}
