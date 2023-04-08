package routes

import (
	"github.com/oxyrinchus/goilerplate/api/controllers"
	"github.com/oxyrinchus/goilerplate/api/middlewares"
	"github.com/oxyrinchus/goilerplate/lib"
)

type MemoRoute struct {
	logger         lib.Logger
	router         lib.Router
	authMiddleware middlewares.AuthMiddleware
	memoController controllers.MemoController
}

// NewMemoRoute initialize memo route
func NewMemoRoute(logger lib.Logger, router lib.Router, authMiddleware middlewares.AuthMiddleware, memoController controllers.MemoController) MemoRoute {
	return MemoRoute{
		logger:         logger,
		router:         router,
		authMiddleware: authMiddleware,
		memoController: memoController,
	}
}

// Setup sets up memo route
func (mr MemoRoute) Setup() {
	mr.logger.Info("Setting up memo route")
	api := mr.router.Gin.Group("/memo").Use(mr.authMiddleware.Handler())
	{
		api.GET("/", mr.memoController.FindAllMemo)
		api.POST("/", mr.memoController.CreateMemo)
		api.GET("/:id", mr.memoController.FindOneMemo)
		api.PUT("/:id", mr.memoController.UpdateMemo)
		api.DELETE("/:id", mr.memoController.DeleteMemo)
	}
}
