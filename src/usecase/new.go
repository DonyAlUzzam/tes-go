package usecase

import (
	"myapp/config"
	"myapp/src/repository"
)

type UsecaseContext struct {
	cfg config.Config
}

type Usecase interface {
	User() UserUsecase
}

func NewUsecase(cfg config.Config) Usecase {
	return &UsecaseContext{
		cfg: cfg,
	}
}

func (c *UsecaseContext) User() UserUsecase {

	repo := repository.NewRepository()
	return NewUserUsecase(repo.User(c.cfg))
}
