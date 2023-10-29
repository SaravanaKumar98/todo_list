package request

import (
	"github.com/go-playground/validator/v10"
)

type Login struct {
	Email    string `validate:"required",json:"email"`
	Password string `validate:"required",json:"password"`
}

func LogIn(email string, password string) error {
	validate := validator.New()
	err := validate.StructExcept(Login{Email: email, Password: password}, "Uuid")
	if err != nil {
		return err
	}
	return nil
}
