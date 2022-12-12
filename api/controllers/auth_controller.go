package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyrinchus/goilerplate/api/controllers/dto"
	"github.com/oxyrinchus/goilerplate/api/controllers/msg"
	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/services"
)

type AuthController struct {
	logger      lib.Logger
	authService services.AuthService
}

// NewAuthController initialize auth controller.
func NewAuthController(logger lib.Logger, authService services.AuthService) AuthController {
	return AuthController{
		logger:      logger,
		authService: authService,
	}
}

// SignUp signs up the user.
func (ac AuthController) SignUp(c *gin.Context) {
	var dto dto.SignUp

	status, data, msg := func() (int, bool, string) {
		if err := c.ShouldBind(&dto); err != nil {
			ac.logger.Error(err)
			return http.StatusBadRequest, false, msg.BAD_REQUEST
		}

		err := ac.authService.SignUp(dto.Email, dto.Password, dto.Name, dto.Role)
		if err != nil {
			return http.StatusInternalServerError, false, msg.CONTACT_SERVER_ADMIN
		}

		return http.StatusOK, true, msg.SIGNUP_SUCCESS
	}()

	ac.logger.Debugf("[SignUp] %+v -> {status:%d, data:%+v, msg:%s}", dto, status, data, msg)
	c.JSON(status, gin.H{"data": data, "msg": msg})
}

// SignIn signs in the user
func (ac AuthController) SignIn(c *gin.Context) {
	var dto dto.SignIn

	status, data, msg := func() (int, bool, string) {
		if err := c.ShouldBind(&dto); err != nil {
			ac.logger.Error(err)
			return http.StatusBadRequest, false, msg.BAD_REQUEST
		}

		accessToken, refreshToken := "", ""

		ok, err := ac.authService.SignIn(dto.Email, dto.Password, &accessToken, &refreshToken)
		if err != nil {
			ac.logger.Error(err)
			return http.StatusInternalServerError, false, msg.CONTACT_SERVER_ADMIN
		}
		if !ok {
			ac.logger.Info("SignIn Fail")
			return http.StatusOK, false, msg.NOT_MATCH_ID_PW
		}

		c.SetCookie(common.ACCESS_TOKEN, accessToken, common.ACCESS_TOKEN_TTL, "/", "localhost", false, true)
		c.SetCookie(common.REFRESH_TOKEN, refreshToken, common.REFRESH_TOKEN_TTL, "/", "localhost", false, true)

		ac.logger.Info("SignIn Success")
		return http.StatusOK, true, msg.SIGNIN_SUCCESS
	}()

	ac.logger.Debugf("[SignIn] %+v -> {status:%d, data:%+v, msg:%s}", dto, status, data, msg)
	c.JSON(status, gin.H{"data": data, "msg": msg})
}
