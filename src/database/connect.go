package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/KaioMarxDEV/gofinance/src/config"
	"github.com/KaioMarxDEV/gofinance/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to Database")
	}

	fmt.Println("Connection opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Auto migration finished")
}
