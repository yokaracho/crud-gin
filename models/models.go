package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Name string `json:"name"`
	Year int    `json:"year"`
}
