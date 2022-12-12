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
		api.GET("/:id", mr.memoController.FindOneMemo)
		api.POST("/create", mr.memoController.CreateMemo)
		api.POST("/update", mr.memoController.UpdateMemo)
		api.GET("/delete/:id", mr.memoController.DeleteMemo)
	}
}
