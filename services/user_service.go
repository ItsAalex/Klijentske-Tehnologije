package services

import (
	"klijentske-tehnologije/models"
	"klijentske-tehnologije/repositories"
	"klijentske-tehnologije/requests"
)

// UserService is an interface that defines the methods a user service should implement.
type UserService interface {
	Create(user requests.CreateUserRequest)
}

// UserServiceImpl is the implementation of the UserService interface.
type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

// NewUserServiceImpl creates a new instance of UserServiceImpl.
func NewUserServiceImpl(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

// Create creates a new user based on the provided CreateUserRequest.
func (u *UserServiceImpl) Create(user requests.CreateUserRequest) {
	// Create a UserModel based on the CreateUserRequest
	userModel := models.User{
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		StudyProgram: user.StudyProgram,
		Index:        user.Index,
		Comment:      user.Comment,
	}

	// Call the UserRepository's Create method to persist the user in the database
	u.UserRepository.Create(userModel)
}
