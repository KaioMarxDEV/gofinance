package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// functin to get env correspondent value from key passed
func Config(key string) string {
	// load .env file and hit errors if does not exists
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error - Loading .env file:\t", err)
	}

	return os.Getenv(key)
}
