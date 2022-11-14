package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/services"
)

type UserController struct {
	userService services.UserService
	logger      lib.Logger
}

func NewUserController(userService services.UserService, logger lib.Logger) UserController {
	return UserController{
		userService: userService,
		logger:      logger,
	}
}

func (uc UserController) GetUserInfo(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		uc.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong user id.",
		})
		return
	}

	user, err := uc.userService.GetUser(uint(id))
	if err != nil {
		uc.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"name":  user.Name,
	})
}
