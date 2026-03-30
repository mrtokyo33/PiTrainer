package usecases

import "github.com/mrtokyo33/PiTrainer/src/repositories"

type GetMeUseCase struct {
	userRepo repositories.IUserRepository
}

func NewGetMeUseCase(repo repositories.IUserRepository) *GetMeUseCase {
	return &GetMeUseCase{userRepo: repo}
}

func (u *GetMeUseCase) Execute(userID uint) (interface{}, error) {
	return u.userRepo.FindByID(userID)
}
