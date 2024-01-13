package migration

import (
	"klijentske-tehnologije/configs"
	userControllers "klijentske-tehnologije/controllers/user"
	"klijentske-tehnologije/repositories"
	"klijentske-tehnologije/routes"
	"klijentske-tehnologije/services"
	//adminControllers"klijentske-tehnologije/controllers/admin"
)

func Serve() {
	// Initialize and connect to the database
	db, err := configs.Connect()
	if err != nil {
		// Handle the error appropriately
		return
	}
	configs.LoadConfig()
	// User Controllers and Services
	userUserController := userControllers.NewUserController(services.NewUserServiceImpl(repositories.NewUserRepositoryImpl(db)))
	userAuthController := userControllers.NewAuthController(services.NewUserServiceImpl(repositories.NewUserRepositoryImpl(db)), db)
	userQuestionController := userControllers.NewQuestionController(services.NewQuestionServiceImpl(repositories.NewQuestionRepositoryImpl(db)))

	// Setup the router
	routes.SetupRouter(
		*userUserController,
		*userAuthController,
		*userQuestionController,
	)
}
