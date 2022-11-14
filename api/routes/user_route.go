package routes

import (
	"github.com/oxyrinchus/goilerplate/api/controllers"
	"github.com/oxyrinchus/goilerplate/lib"
)

type UserRoute struct {
	logger lib.Logger
	router lib.Router
	userController controllers.UserController
}

// NewUserRoute initialize user route
func NewUserRoute(logger lib.Logger, router lib.Router, userController controllers.UserController) UserRoute {
	return UserRoute{
		logger: logger,
		router: router,
		userController: userController,
	}
}

// Setup sets up user route
func (ur UserRoute) Setup() {
	ur.logger.Info("Setting up user route")
	api := ur.router.Gin.Group("/user")
	{
		api.GET("/:id", ur.userController.GetUserInfo)
	}
}
