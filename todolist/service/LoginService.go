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

type Login struct {
	Email    string `validate:"required",json:"email"`
	Password string `validate:"required",json:"password"`
}

func LogIn(email string, password string) (bool, *model.User, error) {
	SignIn := database.Todolist.Collection("signin")
	filter := bson.M{"email": email}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := SignIn.CountDocuments(ctx, filter)
	if err != nil {
		return false, &model.User{}, err
	}

	if count == 0 {
		return false, &model.User{}, errors.New("User not found..")
	}
	var User model.User
	err = SignIn.FindOne(ctx, filter).Decode(&User)
	if err != nil {
		return false, &model.User{}, err
	}

	checkPassword := helper.CompareHash(User.Password, password)

	if checkPassword {
		return true, &User, nil
	}
	return false, &model.User{}, err
}
