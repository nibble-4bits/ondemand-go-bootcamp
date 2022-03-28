package router

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpAPI/v1/controller"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

func getCommentByID(r *gin.RouterGroup, service usecase.CommentService) {
	r.GET("/comments/:id", controller.GetCommentByID(service))
}

func RegisterCommentRoutes(r *gin.RouterGroup, service usecase.CommentService) {
	getCommentByID(r, service)
}
