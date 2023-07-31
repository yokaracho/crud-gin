package controllers

import (
	"crudgin/initializers"
	"crudgin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCar(c *gin.Context) {
	var body struct {
		Name string
		Year int
	}
	c.Bind(&body)

	car := models.Car{Name: body.Name, Year: body.Year}

	result := initializers.DB.Create(&car)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"car": car,
	})
}

func GetCar(c *gin.Context) {
	var cars []models.Car
	initializers.DB.Find(&cars)
	c.JSON(200, gin.H{
		"cars": cars,
	})
}

func GetCarId(c *gin.Context) {
	id := c.Param("id")

	var posts models.Car
	initializers.DB.First(&posts, id)

	c.JSON(200, gin.H{
		"cars": posts,
	})
}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		Year int
	}
	c.Bind(&body)

	var cars models.Car
	initializers.DB.First(&cars, id)

	initializers.DB.Model(&cars).Updates(models.Car{
		Name: body.Name,
		Year: body.Year,
	})

	c.JSON(200, gin.H{
		"cars": cars,
	})
}

func DeleteCar(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Car{}, id)

	c.Status(http.StatusOK)
}
