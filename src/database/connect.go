package database

import (
	"fmt"
	"log"

	"github.com/KaioMarxDEV/gofinance/src/config"
	"github.com/KaioMarxDEV/gofinance/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to Database")
	}

	fmt.Println("Connection opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Auto migration finished")
}
