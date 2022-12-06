package routes

import "go.uber.org/fx"

// exports routes dependency
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoute),
	fx.Provide(NewAuthRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

// create a new routes
func NewRoutes(userRoute UserRoute, authRoute AuthRoute) Routes {
	return Routes{
		userRoute,
		authRoute,
	}
}

// set up routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
