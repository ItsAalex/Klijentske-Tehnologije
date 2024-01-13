package routes

import (
	fmt "fmt"
	adminController "klijentske-tehnologije/controllers/admin"
	userController "klijentske-tehnologije/controllers/user"
	http "net/http"

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	userUserController userController.UserController,
	userQuestionController userController.QuestionController,
	adminUserController adminController.UserController,

) *gin.Engine {

	// Create a new Gin router with default middleware
	r := gin.Default()

	// Add Gin recovery middleware to handle panics gracefully
	r.Use(gin.Recovery())

	// add  swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/admin/user", adminUserController.FindAll)

	r.DELETE("/admin/user/:id", adminUserController.Delete)

	r.POST("/admin/user/approve")

	r.POST("/user/register", userUserController.Create)

	r.POST("/user/question", userQuestionController.Create)

	// Define a default route for the root URL ("/")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Create an HTTP server with the configured router
	server := &http.Server{
		Addr:    ":8080", // Listen on port 8080
		Handler: r,       // Use the Gin router as the request handler
	}

	// Print a message indicating that the server has started
	fmt.Println("Started Server")

	// Start the HTTP server and listen for incoming requests
	server.ListenAndServe()

	// Return the configured Gin router
	return r
}
