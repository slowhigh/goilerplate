package main

import (
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/someday-94/TypeGoMongo-Server/api"
	"github.com/someday-94/TypeGoMongo-Server/controller"
	"github.com/someday-94/TypeGoMongo-Server/docs"
	"github.com/someday-94/TypeGoMongo-Server/middlewares"
	"github.com/someday-94/TypeGoMongo-Server/repository"
	"github.com/someday-94/TypeGoMongo-Server/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	database repository.Database = repository.NewRepository()
	videoRepository repository.VideoRepository = repository.NewVideoRepository(database)

	videoService service.VideoService = service.New(videoRepository)
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Someday - Video API"
	docs.SwaggerInfo.Description = "Someday - Video List API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer database.CloseDB()

	server := gin.Default()

	server.GET("/", middlewares.PlaygroundHandler())
	server.POST("query", middlewares.GraphQLHandler())


	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
