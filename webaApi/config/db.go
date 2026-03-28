package config

import (
	"webapi/webaApi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:tiger@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Employee{})
	return db
}
