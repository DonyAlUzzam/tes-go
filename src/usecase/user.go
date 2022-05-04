package usecase

import (
	"context"
	"myapp/src/model"
	"myapp/src/repository"
	users "myapp/src/usecase/user"
)

type UserUsecaseContext struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	GetOneByEmail(ctx context.Context, params *users.GetOneByEmailRequest) (*model.Users, error)
	GetAllDataUsers(ctx context.Context, params *users.GetAllDataUsersRequest) ([]model.Users, error)
	UpdateDataByEmail(ctx context.Context, params *users.UpdateDataByEmailRequest) (*model.Users, error)
	DeleteUserByEmail(ctx context.Context, params *users.GetOneByEmailRequest) (*model.Users, error)
	Login(ctx context.Context, params *users.LoginRequest) (*model.Users, error)
	CreateUser(ctx context.Context, params *users.CreateUserRequest) (*model.Users, error)
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

func (c *UserUsecaseContext) GetAllDataUsers(ctx context.Context, keywords *users.GetAllDataUsersRequest) ([]model.Users, error) {
	users, err := c.userRepository.GetAllDataUsers(keywords.Keywords)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *UserUsecaseContext) UpdateDataByEmail(ctx context.Context, params *users.UpdateDataByEmailRequest) (*model.Users, error) {
	users, err := c.userRepository.UpdateDataByEmail(params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *UserUsecaseContext) DeleteUserByEmail(ctx context.Context, params *users.GetOneByEmailRequest) (*model.Users, error) {

	user, err := c.userRepository.DeleteUserByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserUsecaseContext) Login(ctx context.Context, params *users.LoginRequest) (*model.Users, error) {
	user, err := c.userRepository.Login(params.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserUsecaseContext) CreateUser(ctx context.Context, params *users.CreateUserRequest) (*model.Users, error) {
	users, err := c.userRepository.CreateUser(params)
	if err != nil {
		return nil, err
	}

	return users, nil
}
