package main

import (
	"net/http"
	"notebook/config"
	"notebook/database"

	"github.com/gin-gonic/gin"
)

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
