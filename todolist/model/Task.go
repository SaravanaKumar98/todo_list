package model

import (
	"context"
	"time"
	"todolist/database"

	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	Name     string    `validator:"required",json:"name",bson:"name"`
	Uuid     string    `bson:"uuid"`
	UserUuid string    `bson:useruuid`
	Date     time.Time `validator:"required",json:"date",bson:"date"`
	Status   string    `validator:"required",json:"status",bson:"status"`
}

func (Task Task) UpdateTask() error {
	TaskCollection := database.Todolist.Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	filter := bson.M{"uuid": Task.Uuid}
	update := bson.M{"$set": bson.M{"status": Task.Status}}
	_, err := TaskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GetTask(UserId string) ([]Task, error) {

	TaskCollection := database.Todolist.Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res, err := TaskCollection.Find(ctx, bson.M{"useruuid": UserId})
	if err != nil {
		return nil, err
	}

	var data []Task

	err = res.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetTodayTask(UserId string) ([]Task, error) {

	TaskCollection := database.Todolist.Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	date := time.Now()
	sdate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	edate := sdate.Add(24 * time.Hour)

	res, err := TaskCollection.Find(ctx, bson.M{"date": bson.M{"$gte": sdate, "$lt": edate.UTC()}})

	if err != nil {
		return nil, err
	}

	var data []Task

	err = res.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
