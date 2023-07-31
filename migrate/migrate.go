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
	initializers.DB.AutoMigrate(&models.Car{})
}
