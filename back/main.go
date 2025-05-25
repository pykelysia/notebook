package main

import (
	"notebook/config"
	"notebook/database"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	database.Open("data.db")

	e := server.Run(config.Load())
	if e != nil {
		panic(e)
	}
}
