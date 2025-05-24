package route

import (
	"notebook/database"

	"github.com/gin-gonic/gin"
)

type Data database.DataModel

func NewRunRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		cxt.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func NewCreateRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		var data Data
		e := cxt.ShouldBindJSON(&data)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "error data",
			})
			return
		}

		e = database.NewDataModel().Create(&database.DataModel{
			UID:  data.UID,
			Data: data.Data,
			Done: data.Done,
		})
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "add failed",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func NewGetRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		var data Data
		e := cxt.ShouldBindJSON(&data)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "error data",
			})
			return
		}

		res, err := database.NewDataModel().Get(data.UID)
		if err != nil {
			cxt.JSON(200, gin.H{
				"message": "this user has no note",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"id":      res.UID,
			"data":    res.Data,
			"done":    res.Done,
			"message": "success",
		})
	}
}
