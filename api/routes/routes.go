package routes

import "go.uber.org/fx"

// exports routes dependency
var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup()
}

type Routes []Route

// create a new routes
func NewRoutes() Routes {
	return Routes{}
}

// set up routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
