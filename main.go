package main

import (
	"crudgin/controllers"
	"crudgin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/car", controllers.CreateCar)
	r.GET("/car", controllers.GetCar)
	r.GET("/car/:id", controllers.GetCarId)
	r.PUT("/car/:id", controllers.UpdateCar)
	r.DELETE("/car/:id", controllers.DeleteCar)
	err := r.Run()
	if err != nil {
		return
	}
}
