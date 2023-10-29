package service

import (
	"context"
	"time"
	"todolist/database"
	"todolist/helper"
	"todolist/model"
)

func TaskAdd(task *model.Task) error {
	taskCollection := database.Todolist.Collection("task")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var err error
	task.Uuid, err = helper.UuidGenerate()
	task.Status = "pending"
	if err != nil {
		return err
	}
	_, err = taskCollection.InsertOne(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
