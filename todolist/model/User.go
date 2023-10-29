package model

type User struct {
	Username string `validate:"required",json:"username"`
	Uuid     string `json:"uuid",bson:"uuid"`
	Email    string `validate:"required,email",json:"email"`
	Password string `validate:"required,min=6",json:"password"`
}
