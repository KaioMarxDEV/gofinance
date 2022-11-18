package useCases

import (
	"github.com/KaioMarxDEV/gofinance/src/User/models"
	"github.com/KaioMarxDEV/gofinance/src/User/repositories"
)

func createUserUseCase(user *models.User) error {

	repositories.Create(user)

	return nil
}
