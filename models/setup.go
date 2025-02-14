package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func ConnectDatabase(name string) {
	var err error
	Client, err = gorm.Open(sqlite.Open(name), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = Client.AutoMigrate(&Book{})
	if err != nil {
		panic(err.Error())
	}
}
