package repositories

import (
	"fmt"
	"klijentske-tehnologije/helpers"
	"klijentske-tehnologije/models"
	"net/smtp"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) *models.User
	FindAll() []models.User
	Delete(userId int)
	SendConfirmationEmail()
}
type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (u *UserRepositoryImpl) Create(user models.User) *models.User {
	result := u.Db.Save(&user)
	helpers.ErrorPanic(result.Error)

	return &user
}

func (u *UserRepositoryImpl) FindAll() []models.User {
	var user []models.User
	result := u.Db.Find(&user)
	helpers.ErrorPanic(result.Error)
	return user
}

func (u *UserRepositoryImpl) Delete(userId int) {
	var user models.User
	result := u.Db.Where("id =?", userId).Delete(&user)
	helpers.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) SendConfirmationEmail() {
	email := "gamerscrazy61@gmail.com"
	password := "zwcd etac undb ddgx"
	auth := smtp.PlainAuth("", email, password, "smtp.gmail.com")

	msg := "Subject: Subject\nMessage"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		[]string{email}, // address of person that we are sending email to.
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
