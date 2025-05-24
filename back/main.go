package main

import (
	"notebook/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	e := server.Run(config.Load())
	if e != nil {
		panic(e)
	}
}
