package route

import (
	"fmt"
	"net/http"
	"notebook/database"
	"notebook/tool"

	"github.com/gin-gonic/gin"
)

func NewRunRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		cxt.HTML(http.StatusOK, "index.html", gin.H{
			"message": "success",
		})
	}
}

func NewCreateRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		data, e := tool.GetData(cxt)
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

func NewDeleteRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		data, e := tool.GetData(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "error data",
			})
			return
		}

		e = database.NewDataModel().Delete(data.UID)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "delete fail",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"message": "delete success",
		})
	}
}

func NewUpdateRoute() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		data, e := tool.GetData(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "sata error",
			})
			return
		}
		e = database.NewDataModel().Updata(&data)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "updata fail",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"message": "updata success",
		})
	}
}

func NewGetRoute() gin.HandlerFunc {
	fmt.Println("123")
	return func(cxt *gin.Context) {
		data, e := tool.GetData(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "error data",
			})
			return
		}

		res, err := database.NewDataModel().Get(data.UID)
		if err != nil {
			cxt.JSON(200, gin.H{
				"message": "this user has no more note",
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
