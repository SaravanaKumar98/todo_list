package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Todolist *mongo.Database

func DBStart() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	Todolist = client.Database("todolist")
	fmt.Println("DB Started...")
}
