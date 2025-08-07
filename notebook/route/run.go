package route

import (
	"fmt"
	"net/http"
	"notebook/database"
	"notebook/tool"
	"strconv"

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
		uad, e := tool.GetUaD(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "error data",
			})
			return
		}

		e = database.NewDataModel().Create(&database.DataModel{
			ID:   database.DataNum,
			Data: uad.Data,
			Done: uad.Done,
		}) //内置一次database.DateNum++
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "add failed",
			})
			return
		}
		__user, _ := database.NewUserModel().Get(uad.UID)
		__user.Code = append(__user.Code, strconv.Itoa(int(database.DataNum-1)))
		e = database.NewUserModel().Updata(&__user)
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

		e = database.NewDataModel().Delete(data.ID)
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

		res, err := database.NewDataModel().Get(data.ID)
		if err != nil {
			cxt.JSON(200, gin.H{
				"message": "this user has no more note",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"id":      res.ID,
			"data":    res.Data,
			"done":    res.Done,
			"message": "success",
		})
	}
}
