package main

import (
	"klijentske-tehnologije/configs"
	repositories "klijentske-tehnologije/repositories"
)

func main() {
	//cmd.Execute()
	db := configs.Connection()

	// Create an instance of UserRepositoryImpl
	userRepository := repositories.NewUserRepositoryImpl(db)

	// Call SendConfirmationEmail on the UserRepositoryImpl instance
	userRepository.SendConfirmationEmail()
}
