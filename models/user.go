package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primarykey"`
	Email     string    `gorm:"coloumn:email;unique" json:"Email"`
	Password  string    `gorm:"coloumn:password" json:"Password"`
	FirstName string    `gorm:"coloumn:firstname" json:"Firstname"`
	LastName  string    `gorm:"coloumn:lastname" json:"Lastname"`
	Address   string    `gorm:"coloumn:address" json:"Address"`
	City      string    `gorm:"coloumn:city" json:"City"`
	Postcode  string    `gorm:"coloumn:post_code" json:"PostCode"`
	Phone     string    `gorm:"coloumn:phone" json:"Phone"`
	CreatedAt time.Time `gorm:"coloumn:created_at" json:"CreatedAt"`
}

func (User) TableName() string {
	return "user"
}
