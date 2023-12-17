package user

import (
	"klijentske-tehnologije/helpers"
	"klijentske-tehnologije/requests"
	"klijentske-tehnologije/responses"
	"klijentske-tehnologije/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserService
	Db          *gorm.DB
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (ctrl *UserController) Create(c *gin.Context) {
	// Logic for creating a user in the database
	createUserRequest := requests.CreateUserRequest{}
	err := c.ShouldBindJSON(&createUserRequest)
	helpers.ErrorPanic(err)

	ctrl.userService.Create(createUserRequest)
	webResponse := responses.Response{
		Code:   http.StatusCreated, // Use StatusCreated for successful resource creation
		Status: "Created",
		Data:   nil,
	}
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}
