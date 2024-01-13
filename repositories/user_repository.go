package repositories

import (
	"klijentske-tehnologije/helpers"
	"klijentske-tehnologije/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) *models.User
	FindAll() []models.User
	Delete(userId int)
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
