package repositories

import (
	"errors"

	"github.com/KaioMarxDEV/gofinance/src/User/models"
)

var database []models.User

func Create(usr *models.User) {
	database = append(database, *usr)
}

func FindByName(name string) (models.User, error) {
	for _, v := range database {
		if v.Name == name {
			return v, nil
		}
	}

	return models.User{}, errors.New("user not found")
}
