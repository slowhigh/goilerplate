package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/someday-94/TypeGoMongo-Server/middlewares"
)

func main() {

	server := gin.Default()

	server.GET("/", middlewares.PlaygroundHandler())
	server.POST("query", middlewares.GraphQLHandler())

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
