package tool

import (
	"notebook/database"

	"github.com/gin-gonic/gin"
)

type UaDModule struct {
	UID    uint          `gorm:"unique;primaryKey" json:"uid"`
	Number uint          `json:"password"`
	Code   database.Code `json:"code"`
	ID     uint          `gorm:"primaryKey;unique" json:"id"`
	Data   string        `json:"data"`
	Done   bool          `json:"done"`
}

func GetData(cxt *gin.Context) (data database.DataModel, e error) {
	e = cxt.ShouldBindJSON(&data)
	return data, e
}

func GetUser(cxt *gin.Context) (user database.UserModel, e error) {
	e = cxt.ShouldBindJSON(&user)
	return user, e
}

func GetUaD(cxt *gin.Context) (uad UaDModule, e error) {
	e = cxt.ShouldBindJSON(&uad)
	return uad, e
}
