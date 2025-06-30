package main

import (
	route "notebook/route/run"

	"github.com/gin-gonic/gin"
)

func BindRoute(server *gin.Engine) {
	server.GET("/", route.NewRunRoute())

	servergroup := server.Group("/")

	servergroup.POST("/add", route.NewCreateRoute())
	servergroup.POST("/delete", route.NewDeleteRoute())
	servergroup.POST("/updata", route.NewUpdateRoute())
	servergroup.POST("/get", route.NewGetRoute())
}
