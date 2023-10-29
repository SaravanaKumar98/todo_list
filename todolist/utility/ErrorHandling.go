package utility

import "github.com/gin-gonic/gin"

func MyError(ctx *gin.Context) {
	msg := ctx.Value("message")
	err := ctx.Value("err")
	ctx.JSON(400, gin.H{"message": msg, "details": err})
	ctx.Abort()
}
