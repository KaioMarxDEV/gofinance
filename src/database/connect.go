package database

import (
	"fmt"
	"log"

	"github.com/KaioMarxDEV/gofinance/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dsn := "host=db port=5432 user=gofinance password=docker dbname=postgres sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to Database")
	}

	fmt.Println("Connection opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Auto migration finished")
}
