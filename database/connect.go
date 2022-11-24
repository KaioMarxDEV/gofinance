package database

import (
	"fmt"
	"log"

	"github.com/KaioMarxDEV/gofinance/config"
	"github.com/KaioMarxDEV/gofinance/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() error {
	var err error

	// Get .env values and format to string connection by gorm format
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to Database")
		return err
	}

	fmt.Println("Connection opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Auto migration finished")
	return nil
}
