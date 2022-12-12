package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyrinchus/goilerplate/api/controllers/msg"
	"github.com/oxyrinchus/goilerplate/common"
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
	status, data, msg := func() (int, map[string]any, string) {
		userID, exists := c.Get(common.CURRENT_USER_ID)
		if !exists {
			uc.logger.Error(common.ERR_EMPTY_USER)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		user, err := uc.userService.GetUserInfoByID(userID.(string))
		if err != nil {
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		return http.StatusOK, map[string]any{"email": user.Email, "name": user.Name}, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}

func (uc UserController) ExistsEmail(c *gin.Context) {
	status, data, msg := func() (int, bool, string) {
		email := c.Param("email")
		if email == "" {
			uc.logger.Error(common.ERR_EMPTY_PARAM)
			return http.StatusBadRequest, false, msg.BAD_REQUEST
		}

		ok, err := uc.userService.ExistsEmail(email)
		if err != nil {
			return http.StatusInternalServerError, false, msg.CONTACT_SERVER_ADMIN
		}

		return http.StatusOK, ok, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}
