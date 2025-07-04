package main

import (
	"net/http"
	"notebook/config"
	"notebook/database"

	"notebook/route"

	"github.com/gin-gonic/gin"
)

func BindRoute(server *gin.Engine) {
	//主界面
	server.GET("/", route.NewRunRoute())

	servergroup := server.Group("/")

	servergroup.POST("/add", route.NewCreateRoute())
	servergroup.POST("/delete", route.NewDeleteRoute())
	servergroup.POST("/updata", route.NewUpdateRoute())
	servergroup.POST("/get", route.NewGetRoute())

	//用户相关
	server.GET("/user", route.RunUser())

	usergroup := server.Group("/user")

	usergroup.POST("/get", route.GetCode())
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
