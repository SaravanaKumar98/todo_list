package service

import (
	"context"
	"errors"
	"time"
	"todolist/database"
	"todolist/helper"
	"todolist/model"

	"go.mongodb.org/mongo-driver/bson"
)

func SingIn(User *model.User) (bool, error) {
	signin := database.Todolist.Collection("signin")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	filter := bson.M{"email": User.Email}
	count, err := signin.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	User.Uuid, err = helper.UuidGenerate()
	if err != nil {
		return false, err
	}

	User.Password = helper.MakeHass(User.Password)

	// fmt.Println(count)

	if count == 0 {
		_, err := signin.InsertOne(ctx, User)
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("Given email already exists..")
	}

	return true, nil
}
