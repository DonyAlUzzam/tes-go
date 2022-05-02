package usecase

import (
	"context"
	"myapp/model"
	"myapp/src/repository"
	users "myapp/src/usecase/user"
)

type UserUsecaseContext struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	GetOneByEmail(ctx context.Context, params *users.GetOneByEmailRequest) (*model.Users, error)
}

func NewUserUsecase(
	userRepository repository.UserRepository,
) UserUsecase {
	return &UserUsecaseContext{
		userRepository: userRepository,
	}
}

func (c *UserUsecaseContext) GetOneByEmail(ctx context.Context, params *users.GetOneByEmailRequest) (*model.Users, error) {
	user, err := c.userRepository.GetOneByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
