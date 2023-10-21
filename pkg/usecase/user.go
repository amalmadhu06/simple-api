package usecase

import (
	"context"

	"github.com/amalmadhu06/simple-api/pkg/repository/interfaces"
	services "github.com/amalmadhu06/simple-api/pkg/usecase/interfaces"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (uc *userUseCase) FindAll(ctx context.Context) (string, error) {
	return uc.userRepo.FindAll(ctx)
}
