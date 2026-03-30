package usecases

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/mrtokyo33/PiTrainer/src/repositories"
	"github.com/mrtokyo33/PiTrainer/src/utils"
)

type LoginUserUseCase struct {
	userRepo repositories.IUserRepository
}

func NewLoginUserUseCase(repo repositories.IUserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{userRepo: repo}
}

func (u *LoginUserUseCase) Execute(username, password string) (string, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
