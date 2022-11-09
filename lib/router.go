package lib

import "github.com/gin-gonic/gin"

type Router struct {
	Gin *gin.Engine
}

// create a new router
func NewRouter() Router {
	engine := gin.New()
	return Router{Gin: engine}
}
