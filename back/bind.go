package main

import (
	route "notebook/route/run"

	"github.com/gin-gonic/gin"
)

func BindRoute(server *gin.Engine) {
	servergroup := server.Group("/")

	servergroup.GET("/test", route.NewRunRoute())
	servergroup.POST("/add", route.NewCreateRoute())
	servergroup.POST("/get", route.NewGetRoute())
}
