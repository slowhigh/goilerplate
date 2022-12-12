package controllers

import "go.uber.org/fx"

// exports controllers dependency
var Module = fx.Options(
	fx.Provide(NewAuthController),
	fx.Provide(NewUserController),
	fx.Provide(NewMemoController),
)
