package controller

import (
	"fmt"
	"todolist/helper"
	"todolist/request"
	"todolist/service"
	"todolist/utility"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `validate:"required",json:"email"`
	Password string `validate:"required",json:"password"`
}

func LogIn(c *gin.Context) {
	login := Login{}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.Set("message", "Binding issue")
		c.Set("details", err.Error())
		utility.MyError(c)
		return
	}

	err = request.LogIn(login.Email, login.Password)
	if err != nil {
		c.Set("message", "validate issue")
		c.Set("details", err.Error())
		utility.MyError(c)
		return
	}

	service, user, err := service.LogIn(login.Email, login.Password)
	if err != nil {
		c.Set("message", "service issue")
		c.Set("details", err.Error())
		utility.MyError(c)
		return
	}
	err = helper.SetAuthCookie(c, user.Uuid)
	if err != nil {
		c.Set("message", "set cookie issue")
		c.Set("details", err.Error())
		fmt.Println(err)
		utility.MyError(c)
		return

	}
	if service {
		c.JSON(200, gin.H{"status code": 200, "message": "Login successfully"})
	}

}
