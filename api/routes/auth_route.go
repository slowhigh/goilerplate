package routes

import (
	"github.com/oxyrinchus/goilerplate/api/controllers"
	"github.com/oxyrinchus/goilerplate/lib"
)

type AuthRoute struct {
	logger lib.Logger
	router lib.Router
	authController controllers.AuthController
}

// NewAuthRoute initialize auth route
func NewAuthRoute(logger lib.Logger, router lib.Router, authController controllers.AuthController) AuthRoute {
	return AuthRoute{
		logger: logger,
		router: router,
		authController: authController,
	}
}

// Setup sets up auth route
func (au AuthRoute) Setup() {
	au.logger.Info("Setting up auth route")
	api := au.router.Gin.Group("/auth")
	{
		api.POST("/signup", au.authController.SignUp)
		api.POST("/signin", au.authController.SignIn)
		
	}
}