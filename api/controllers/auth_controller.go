package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyrinchus/goilerplate/api/controllers/dto"
	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/services"
)

type AuthController struct {
	logger      lib.Logger
	authService services.AuthService
}

func NewAuthController(logger lib.Logger, authService services.AuthService) AuthController {
	return AuthController{
		logger:      logger,
		authService: authService,
	}
}

func (ac AuthController) SignUp(c *gin.Context) {
	var dto dto.SignUp

	if err := c.ShouldBind(&dto); err != nil {
		ac.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}

	result, err := ac.authService.SignUp(dto.Email, dto.Password, dto.Name, dto.Role)
	if err != nil {
		ac.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (ac AuthController) SignIn(c *gin.Context) {
	var dto dto.SignIn

	if err := c.ShouldBind(&dto); err != nil {
		ac.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}

	accessToken, refreshToken, err := ac.authService.SignIn(dto.Email, dto.Password)
	if err != nil {
		ac.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false})
		return
	}

	c.SetCookie(common.ACCESS_TOKEN, accessToken, common.ACCESS_TOKEN_TTL, "/", "localhost", false, true)
	c.SetCookie(common.REFRESH_TOKEN, refreshToken, common.REFRESH_TOKEN_TTL, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"result": true})
}