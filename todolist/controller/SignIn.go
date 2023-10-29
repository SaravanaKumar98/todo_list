package controller

import (
	"todolist/model"
	"todolist/request"
	"todolist/service"
	"todolist/utility"

	"github.com/gin-gonic/gin"
)

func SignIN(c *gin.Context) {

	User := model.User{}
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.Set("message", "SignIn validate issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	validate, err := request.SignIn(&User)

	if err != nil {
		c.Set("message", "SignIn validate issue")
		c.Set("err", err.Error())
		utility.MyError(c)
	}

	if validate {
		service, err := service.SingIn(&User)
		if err != nil {
			c.Set("message", "SignIn Service issue")
			c.Set("err", err.Error())
			utility.MyError(c)
		}
		if service {
			c.JSON(200, gin.H{"status": 200, "message": "User created successfull."})
		}
	}

	// c.JSON(200, gin.H{"user": user})
}
