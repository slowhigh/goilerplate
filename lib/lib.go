package lib

import "go.uber.org/fx"

// exports libraries dependency
var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
)