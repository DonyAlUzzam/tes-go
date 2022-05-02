package repository

import (
	"myapp/config"
	"myapp/model"
)

type UserRepositoryContext struct {
	cfg config.Config
}

type UserRepository interface {
	GetOneByEmail(email string) (*model.Users, error)
}

func NewUserRepository(cfg config.Config) UserRepository {
	return &UserRepositoryContext{
		cfg: cfg,
	}
}

func (c *UserRepositoryContext) GetOneByEmail(email string) (*model.Users, error) {
	user := &model.Users{}
	if err := c.cfg.Slave().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
