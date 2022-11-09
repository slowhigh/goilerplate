package middlewares

import "go.uber.org/fx"

// exports middlewares dependency
var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

// create a new middlewares
func NewMiddlewares(
	corsMiddleware CorsMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
	}
}

// sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
