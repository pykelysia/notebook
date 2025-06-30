package tool

import (
	"notebook/database"

	"github.com/gin-gonic/gin"
)

func GetData(cxt *gin.Context) (data database.DataModel, e error) {
	e = cxt.ShouldBindJSON(&data)
	return data, e
}
