package controller

import (
	"fmt"
	"time"
	"todolist/model"
	"todolist/request"
	"todolist/service"
	"todolist/utility"

	"github.com/gin-gonic/gin"
)

func TaskAdd(c *gin.Context) {
	TaskReq := request.TaskRequest{}
	err := c.ShouldBindJSON(&TaskReq)
	if err != nil {
		c.Set("message", "Binding issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	err = request.TaskAdd(&TaskReq)
	if err != nil {
		c.Set("message", "validate issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}
	Task := model.Task{}
	Date, err := time.Parse("2006-01-02 15:04:05", TaskReq.Date)
	if err != nil {
		c.Set("message", "Date Binding issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return

	}
	Task.Date = Date.UTC()
	Task.Name = TaskReq.Name
	Task.Status = TaskReq.Status
	userUuid := c.Value("AuthId")
	Task.UserUuid = fmt.Sprintf("%v", userUuid)

	if userUuid == nil {
		if err != nil {
			c.Set("message", "service issue")
			c.Set("err", err.Error())
			utility.MyError(c)
			return
		}
	}
	err = service.TaskAdd(&Task)

	if err != nil {
		c.Set("message", "service issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	c.JSON(200, gin.H{"status code": 200, "message": "Task added successfully..!"})

}

func TaskUpdate(c *gin.Context) {
	Task := model.Task{}
	err := c.ShouldBindJSON(&Task)
	if err != nil {
		c.Set("message", "Binding issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	// err = request.TaskAdd(&Task)
	// if err != nil {
	// 	c.Set("message", "validate issue")
	// 	c.Set("err", err.Error())
	// 	utility.MyError(c)
	// 	return
	// }

	err = (Task).UpdateTask()

	if err != nil {
		c.Set("message", "Update issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	c.JSON(200, gin.H{"status code": 200, "message": "Task Update successfully..!"})

}

func TaskGet(c *gin.Context) {

	userId := c.Value("AuthId")

	if userId == nil {

		c.Set("message", "controller issue")
		c.Set("err", "user id not funt ")
		utility.MyError(c)
		return

	}
	UserUuid := fmt.Sprintf("%v", userId)
	tasks, err := model.GetTask(UserUuid)

	if err != nil {
		c.Set("message", "get task model issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	c.JSON(200, gin.H{"status code": 200, "data": tasks})

}

func TaskGetToday(c *gin.Context) {

	userId := c.Value("AuthId")

	if userId == nil {

		c.Set("message", "controller issue")
		c.Set("err", "user id not funt ")
		utility.MyError(c)
		return

	}
	UserUuid := fmt.Sprintf("%v", userId)
	tasks, err := model.GetTodayTask(UserUuid)

	if err != nil {
		c.Set("message", "get task model issue")
		c.Set("err", err.Error())
		utility.MyError(c)
		return
	}

	c.JSON(200, gin.H{"status code": 200, "data": tasks})

}
