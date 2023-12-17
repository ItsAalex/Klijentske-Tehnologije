package repositories

import (
	"klijentske-tehnologije/helpers"
	"klijentske-tehnologije/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) *models.User
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
