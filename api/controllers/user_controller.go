package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/services"
)

var ErrNotFoundUser = errors.New("user: not found user")

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
	userID, exists := c.Get(common.CURRENT_USER_ID)
	if !exists {
		uc.logger.Error("not exists the current user's ID")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	user, err := uc.userService.FindUserByID(userID.(string))
	if err != nil {
		uc.logger.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"name":  user.Name,
	})
}
