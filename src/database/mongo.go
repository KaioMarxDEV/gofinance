package database

import (
	"fmt"

	"github.com/gofiber/storage/mongodb"
)

func Init() *mongodb.Storage {
	return mongodb.New(mongodb.Config{
		ConnectionURI: fmt.Sprint("mongodb+srv://gofinance:Ke6_b2LKZkJEKh.@cluster0.mmskzgy.mongodb.net/?retryWrites=true&w=majority"),
		Reset:         false,
	})
}
