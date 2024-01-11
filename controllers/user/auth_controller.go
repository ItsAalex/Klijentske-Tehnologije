package user

import (
	configs "klijentske-tehnologije/configs"
	models "klijentske-tehnologije/models"
	services "klijentske-tehnologije/services"
	http "net/http"

	gin "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	var body struct {
		Firstname string
		Lastname  string
		Email     string
		Password  string
		Address   string
		City      string
		Postcode  string
		Phone     string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create user with hashed password
	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Address:   body.Address,
		City:      body.City,
		Postcode:  body.Postcode,
		Phone:     body.Phone,
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
