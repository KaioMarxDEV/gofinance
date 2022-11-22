package repositories

import (
	"time"

	"github.com/KaioMarxDEV/gofinance/src/User/models"
	"github.com/KaioMarxDEV/gofinance/src/database"
)

func Create(usr *models.User) {
	store := database.Init()

	store.Set("username", []byte(usr.Name), time.Duration(time.Now().Unix()))
}
