package main

import (
	"crudgin/initializers"
	"crudgin/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Car{})
	if err != nil {
		return
	}
}
