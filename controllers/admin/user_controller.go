package admin

import (
	responses "klijentske-tehnologije/responses"
	services "klijentske-tehnologije/services"
	http "net/http"

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
