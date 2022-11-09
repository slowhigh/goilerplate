package bootstrap

import (
	"github.com/oxyrinchus/goilerplate/api/controllers"
	"github.com/oxyrinchus/goilerplate/api/middlewares"
	"github.com/oxyrinchus/goilerplate/api/routes"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/repositories"
	"github.com/oxyrinchus/goilerplate/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repositories.Module,
)
