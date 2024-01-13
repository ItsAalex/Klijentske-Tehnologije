package services

import (
	"klijentske-tehnologije/models"
	"klijentske-tehnologije/repositories"
	"klijentske-tehnologije/requests"
	"klijentske-tehnologije/responses"
)

// UserService is an interface that defines the methods a user service should implement.
type UserService interface {
	Create(user requests.CreateUserRequest)
	FindAll() []responses.UserResponse
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

// FindAll retrieves a list of users and converts them into UserResponse objects.
func (u *UserServiceImpl) FindAll() []responses.UserResponse {
	// Call the UserRepository's FindAll method to retrieve a list of users
	result := u.UserRepository.FindAll()

	// Initialize an empty slice to store UserResponse objects
	var users []responses.UserResponse

	// Convert UserModel instances to UserResponse instances
	for _, value := range result {
		user := responses.UserResponse{
			ID:           value.ID,
			Firstname:    value.Firstname,
			Lastname:     value.Lastname,
			Email:        value.Email,
			StudyProgram: value.StudyProgram,
			Index:        value.Index,
			Comment:      value.Comment,
		}
		users = append(users, user)
	}

	// Return the list of UserResponse objects
	return users
}
