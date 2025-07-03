package main

import (
	"net/http"
	"notebook/config"
	"notebook/database"

	"notebook/route"

	"github.com/gin-gonic/gin"
)

func BindRoute(server *gin.Engine) {
	server.GET("/", route.NewRunRoute())

	servergroup := server.Group("/")

	servergroup.POST("/add", route.NewCreateRoute())
	servergroup.POST("/delete", route.NewDeleteRoute())
	servergroup.POST("/updata", route.NewUpdateRoute())
	servergroup.POST("/get", route.NewGetRoute())

	usergroup := server.Group("/user")

	usergroup.POST("/signup", route.SignUp())
	usergroup.POST("/signin", route.SignIn())
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.StaticFS("/static", http.Dir("./static"))
	BindRoute(server)

	database.Open("data.db")

	e := server.Run(config.Load())
	if e != nil {
		panic(e)
	}
}
