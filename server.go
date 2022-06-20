package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/someday-94/TypeGoMongo-Server/controller"
	"github.com/someday-94/TypeGoMongo-Server/middlewares"
	"github.com/someday-94/TypeGoMongo-Server/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	// Case-1
	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	//apiRoutes := server.Group("/api")

	// Case-2
	//server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	//apiRoutes := server.Group("/api", middlewares.BasicAuth())

	// Case-1 처럼하게 되면 아래 /view 경로일때도 middlewares.BasicAuth() 가 동작하여 id, pw를 입력해야 하지만,
	// Case-2 처럼하게 되면 /api 일때만 middlewares.BasicAuth()가 동작하고
	// 기존의 gin.Recovery(), middlewares.Logger(), gindump.Dump()는 모든 path("/api", "/view" 등)에 동일하게 적용된다.

	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	// region 프론트 앤드 영역 / 나중에 지우기
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")
	// endregion

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// server.Run(":8080")
	// 나중에 docker-compose.yaml 파일 또는 Dockerfile 파일에 PORT 변수 값을 지정해 주면 그걸 가져다 사용할 수도 있을거 같다.
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
