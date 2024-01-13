package admin

import (
	"klijentske-tehnologije/helpers"
	responses "klijentske-tehnologije/responses"
	services "klijentske-tehnologije/services"
	http "net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
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

func (ctrl *UserController) FindAll(c *gin.Context) {
	// Logic for retrieving all users from the database
	users := ctrl.userService.FindAll()
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: gin.H{
			"data": users,
		},
	}
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (ctrl *UserController) Delete(c *gin.Context) {
	// Logic for deleting a user from the database by ID
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	helpers.ErrorPanic(err)
	ctrl.userService.Delete(id)
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
