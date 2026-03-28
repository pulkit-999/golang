package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Id         int    `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}
