package lib

import "go.uber.org/fx"

// exports libraries dependency
var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
	fx.Provide(NewRouter),
)
