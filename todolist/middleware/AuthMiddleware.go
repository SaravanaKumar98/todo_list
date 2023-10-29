package middleware

import (
	"context"
	"time"
	"todolist/database"
	"todolist/helper"
	"todolist/utility"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, err := ctx.Cookie("auth")
		if err != nil {
			ctx.Set("message", "Authentication cookie Issue..")
			utility.MyError(ctx)
			return
		}

		Uuid, checkUser, err := checkuser(uuid)

		if err != nil {
			ctx.Set("message", "User not found Issue..")
			utility.MyError(ctx)
			return
		}

		if checkUser {
			ctx.Set("AuthId", Uuid)
			ctx.Next()
			return
		}

	}
}

func checkuser(uuid string) (string, bool, error) {

	user := database.Todolist.Collection("signin")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	Uuid, err := helper.Crypt(uuid, "decrypt")
	if err != nil {
		return "", false, err
	}

	count, err := user.CountDocuments(ctx, bson.M{"uuid": Uuid})

	if err != nil {
		return "", false, err
	}
	if count > 0 {
		return Uuid, true, nil
	}
	return "", false, nil
}
