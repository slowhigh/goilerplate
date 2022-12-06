package commands

import (
	"os"

	"github.com/oxyrinchus/goilerplate/api/middlewares"
	"github.com/oxyrinchus/goilerplate/api/routes"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/spf13/cobra"
)

type ServeCommand struct{}

// create a new serve command
func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}

func (sc *ServeCommand) Short() string {
	return "serve app"
}

func (sc *ServeCommand) Setup(cmd *cobra.Command) {}

func (sc *ServeCommand) Run() lib.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		env lib.Env,
		router lib.Router,
		route routes.Routes,
		logger lib.Logger,
	) {
		middleware.Setup()
		route.Setup()

		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Gin.Run(":" + os.Getenv("PORT"))
		} else {
			_ = router.Gin.Run(":" + env.ServerPort)
		}
	}
}
