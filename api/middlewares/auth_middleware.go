package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/services"
)

type AuthMiddleware struct {
	logger      lib.Logger
	authService services.AuthService
}

func NewAuthMiddleware(logger lib.Logger, authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{
		logger:      logger,
		authService: authService,
	}
}

func (am AuthMiddleware) Setup() {}

func (am AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, _ := c.Cookie(common.ACCESS_TOKEN)
		refreshToken, _ := c.Cookie(common.REFRESH_TOKEN)
		
		am.logger.Debug("accessToken : ", accessToken)
		am.logger.Debug("refreshToken : ", refreshToken)
		
		userID, newAccessToken, err := am.authService.Authorize(accessToken, refreshToken)
		if err != nil {
			am.logger.Info(err)
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		if newAccessToken != "" {
			am.logger.Debug("newAccessToken : ", newAccessToken)
			c.SetCookie(common.ACCESS_TOKEN, newAccessToken, common.ACCESS_TOKEN_TTL, "/", "localhost", false, true)
		}
		
		am.logger.Debug("userID : ", userID)
		c.Set(common.CURRENT_USER_ID,  userID)

		c.Next()
	}
}