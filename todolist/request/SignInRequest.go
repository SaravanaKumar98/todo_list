package request

import (
	"todolist/model"

	"github.com/go-playground/validator/v10"
)

func SignIn(User *model.User) (bool, error) {

	validate := validator.New()

	err := validate.StructExcept(User, "Uuid")

	if err != nil {
		return false, err
	}

	return true, nil

}
