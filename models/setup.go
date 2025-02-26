package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func ConnectDatabase() {
	var err error
	dbUrl := "postgres://admin:admin@0.0.0.0:5432/db"
	Client, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = Client.AutoMigrate(&Book{})
	if err != nil {
		panic(err.Error())
	}
}
