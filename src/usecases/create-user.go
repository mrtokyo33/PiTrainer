package usecases

import (
	"github.com/mrtokyo33/PiTrainer/src/models"
	"github.com/mrtokyo33/PiTrainer/src/repositories"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := models.NewUser(username, string(hashedPassword))
	if err != nil {
		return err
	}

	return u.userRepo.Save(user)
}
