package repository

import "myapp/config"

type RepositoryContext struct {
}

type Repository interface {
	User(cfg config.Config) UserRepository
}

func NewRepository() Repository {
	return &RepositoryContext{}
}

func (c *RepositoryContext) User(cfg config.Config) UserRepository {
	return NewUserRepository(cfg)
}
