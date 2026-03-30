package usecases

import (
	"github.com/mrtokyo33/PiTrainer/src/models"
	"github.com/mrtokyo33/PiTrainer/src/repositories"
)

type CreateUserUseCase struct {
	userRepo repositories.IUserRepository
}

func NewCreateUserUseCase(userRepo repositories.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo: userRepo,
	}
}

func (u *CreateUserUseCase) Execute(username, password string) error {
	user, err := models.NewUser(username, password)
	if err != nil {
		return err
	}

	return u.userRepo.Save(user)
}
