package router

import (
	"todolist/controller"
	"todolist/middleware"

	"github.com/gin-gonic/gin"
)

func RouteStart() {
	app := gin.New()
	user := app.Group("user") //AuthMiddleware
	user.Use(middleware.AuthMiddleware())

	// app.GET("/aa", func(ctx *gin.Context) {
	// 	ctx.SetCookie("auth", "abc", 3600, "/", "localhost", false, true)
	// 	ctx.JSON(200, gin.H{"a": "sss"})
	// })
	app.POST("/signin", controller.SignIN)
	app.POST("/login", controller.LogIn)
	user.POST("/task/add", controller.TaskAdd)
	user.POST("/task/update", controller.TaskUpdate)
	user.GET("/task/get", controller.TaskGet)
	user.GET("/task/get_today", controller.TaskGetToday)
	app.POST("/sign", controller.SignIN)
	app.Run(":8888")

}
