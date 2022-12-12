package services

import "go.uber.org/fx"

// exports services dependency
var Module = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewUserService),
	fx.Provide(NewMemoService),
)
