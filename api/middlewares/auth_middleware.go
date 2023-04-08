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

// NewAuthMiddleware initialize auth middleware.
func NewAuthMiddleware(logger lib.Logger, authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{
		logger:      logger,
		authService: authService,
	}
}

// Setup sets up auth middleware.
func (am AuthMiddleware) Setup() {}

// Handler handles middleware functionality.
func (am AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, _ := c.Cookie(common.ACCESS_TOKEN)
		refreshToken, _ := c.Cookie(common.REFRESH_TOKEN)
		
		am.logger.Debug("accessToken : ", accessToken)
		am.logger.Debug("refreshToken : ", refreshToken)
		
		userID, newAccessToken, err := am.authService.Authorize(accessToken, refreshToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			c.Abort()
			return
		} else if userID == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		if newAccessToken != "" {
			am.logger.Debug("newAccessToken : ", newAccessToken)
			c.SetCookie(common.ACCESS_TOKEN, newAccessToken, common.ACCESS_TOKEN_TTL, "/", "", false, true)
		}
		
		am.logger.Debug("userID : ", userID)
		c.Set(common.CURRENT_USER_ID,  userID)

		c.Next()
	}
}