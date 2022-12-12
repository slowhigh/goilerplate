package middlewares

import (
	"github.com/oxyrinchus/goilerplate/lib"
	cors "github.com/rs/cors/wrapper/gin"
)

type CorsMiddleware struct {
	router lib.Router
	logger lib.Logger
	env    lib.Env
}

// NewCorsMiddleware initialize cors middleware.
func NewCorsMiddleware(router lib.Router, logger lib.Logger, env lib.Env) CorsMiddleware {
	return CorsMiddleware{
		router: router,
		logger: logger,
		env:    env,
	}
}

// Setup sets up cors middleware.
func (cm CorsMiddleware) Setup() {
	cm.logger.Info("Setting up cors middleware")

	cm.router.Gin.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		Debug:            cm.env.Environment == "development",
	}))
}
