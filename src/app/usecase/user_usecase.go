package usecase

import (
	"context"
	"pMade.api/src/app/domain/model"
	"pMade.api/src/app/domain/repository"
)

type UserUseCase interface {
	GetUser(ctx context.Context) ([]*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func (uu userUseCase) GetUser(ctx context.Context) (users []*model.User, err error) {
	users, err = uu.userRepository.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}
