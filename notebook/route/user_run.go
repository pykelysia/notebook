package route

import (
	"fmt"
	"notebook/database"
	"notebook/tool"

	"github.com/gin-gonic/gin"
)

func RunUser() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		cxt.HTML(200, "user.html", gin.H{
			"message": "success",
		})
	}
}

func SignIn() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		user, e := tool.GetUser(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "error data",
			})
			return
		}

		__user, err := database.NewUserModel().Get(user.UID)
		if err != nil {
			cxt.JSON(200, gin.H{
				"message": "this user disappear",
			})
			return
		}
		if user.Number != __user.Number {
			cxt.JSON(200, gin.H{
				"messsage": "password error",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"codemax": len(__user.Code),
			"code":    __user.Code,
			"message": "success",
		})
	}
}

func SignUp() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		user, e := tool.GetUser(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "data error",
			})
			return
		}

		_, err := database.NewUserModel().Get(user.UID)
		if err == nil {
			cxt.JSON(200, gin.H{
				"message": "user has exist",
			})
			return
		}
		e = database.NewUserModel().Create(&database.UserModel{
			UID:    user.UID,
			Number: user.Number,
		})
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "add user failed",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"messsge": "success",
		})
	}
}

func GetCode() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		user, e := tool.GetUser(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "data error",
			})
			return
		}

		__user, _ := database.NewUserModel().Get(user.UID)
		if __user.UID == 0 {
			cxt.JSON(200, gin.H{
				"message": "user has not exist",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"code":    __user.Code,
			"message": "success",
		})
	}
}

func CheckUID() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		user, e := tool.GetUser(cxt)
		if e != nil {
			cxt.JSON(200, gin.H{
				"message": "data error",
			})
			return
		}

		__user, _ := database.NewUserModel().Get(user.UID)
		fmt.Println(__user)
		fmt.Println(user)
		if __user.UID == 0 {
			fmt.Println("?")
			cxt.JSON(200, gin.H{
				"message": "user has not exist",
			})
			return
		}
		cxt.JSON(200, gin.H{
			"message": "right",
		})
	}
}
