package request

import (
	"github.com/go-playground/validator/v10"
)

type TaskRequest struct {
	Name     string `validator:"required",json:"name",bson:"name"`
	Uuid     string `bson:"uuid"`
	UserUuid string `bson:useruuid`
	Date     string `validator:"required",json:"date",bson:"date"`
	Status   string `validator:"required",json:"status",bson:"status"`
}

func TaskAdd(task *TaskRequest) error {
	validate := validator.New()
	err := validate.StructExcept(task, "Uudid", "UserUuid")
	if err != nil {
		return err
	}
	return nil
}
