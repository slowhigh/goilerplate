package middlewares

import "go.uber.org/fx"

// exports middlewares dependency
var Module = fx.Options(
	fx.Provide(NewMiddlewares),
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewAuthMiddleware),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

// create a new middlewares
func NewMiddlewares(
	corsMiddleware CorsMiddleware,
	authMiddleware AuthMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
		authMiddleware,
	}
}

// sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
