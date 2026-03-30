package repositories

import "github.com/mrtokyo33/PiTrainer/src/models"

type IUserRepository interface {
	Save(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}
