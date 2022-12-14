package database

import (
	"fmt"
	"log"

	"github.com/KaioMarxDEV/gofinance/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() bool {
	var (
		err error
		dsn string
	)

	// host=db when using with docker compose
	// host=localhost when using with air live reload
	dsn = "host=db  port=5432 user=gofinance password=docker dbname=postgres sslmode=disable"

	// open the connection with database using ORM library named gorm and postgres driver
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if connection fails panic the system
	if err != nil {
		log.Panic("Failed to connect to Database")
	}

	fmt.Println("Connection opened to Database")
	DB.AutoMigrate(&models.User{}, &models.Transaction{}) // migrate the models to database creating tables automatically
	fmt.Println("Auto migration finished")
	return true
}
