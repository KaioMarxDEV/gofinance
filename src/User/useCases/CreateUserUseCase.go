package useCases

import (
	"github.com/KaioMarxDEV/gofinance/src/User/models"
	repositories "github.com/KaioMarxDEV/gofinance/src/User/repository"
)

func createUserUseCase(user *models.User) error {

	repositories.Create(user)

	return nil
}
