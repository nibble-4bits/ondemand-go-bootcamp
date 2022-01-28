package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func CreateRouterGroup(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/v1")

	return v1
}

// StartServer makes the HTTP server start listening for requests
func StartServer(router *gin.Engine) {
	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
