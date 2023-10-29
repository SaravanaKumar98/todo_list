package main

import (
	"todolist/database"
	"todolist/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// helper.Mycrypt()

	database.DBStart()
	router.RouteStart()

}
