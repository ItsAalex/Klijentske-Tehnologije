package user

import (
	configs "klijentske-tehnologije/configs"
	models "klijentske-tehnologije/models"
	requests "klijentske-tehnologije/requests"
	services "klijentske-tehnologije/services"
	http "net/http"

	gin "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	userService services.UserService
	Db          *gorm.DB
}

func NewAuthController(service services.UserService) *AuthController {
	return &AuthController{
		userService: service,
	}
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	db := configs.Connection()

	var request requests.CreateUserRequest
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hashedPassword, err := services.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create user with hashed password
	user := models.User{
		Email:     request.Email,
		Password:  hashedPassword,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Address:   request.Address,
		City:      request.City,
		Postcode:  request.Postcode,
		Phone:     request.Phone,
	}
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
