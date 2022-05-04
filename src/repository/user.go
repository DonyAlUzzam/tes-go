package repository

import (
	"myapp/config"
	"myapp/src/model"
	users "myapp/src/usecase/user"
)

type UserRepositoryContext struct {
	cfg config.Config
}

type UserRepository interface {
	GetOneByEmail(email string) (*model.Users, error)
	GetAllDataUsers(keywords string) ([]model.Users, error)
	UpdateDataByEmail(data *users.UpdateDataByEmailRequest) (*model.Users, error)
	DeleteUserByEmail(email string) (*model.Users, error)
	Login(email string) (*model.Users, error)
	CreateUser(data *users.CreateUserRequest) (*model.Users, error)
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

func (c *UserRepositoryContext) GetAllDataUsers(keywords string) ([]model.Users, error) {
	// user := &model.Users{}
	var users []model.Users
	if err := c.cfg.Slave().Where("email LIKE ? OR nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (c *UserRepositoryContext) UpdateDataByEmail(data *users.UpdateDataByEmailRequest) (*model.Users, error) {
	user := &model.Users{}
	user.Email = data.Email
	user.Alamat = data.Alamat
	user.Ktp = data.Ktp
	user.Nama = data.Nama
	user.NoHp = data.NoHp
	user.Password = data.Password

	if err := c.cfg.Slave().Where("email = ?", data.Email).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserRepositoryContext) DeleteUserByEmail(email string) (*model.Users, error) {
	user := &model.Users{}
	if err := c.cfg.Slave().Where("email = ?", email).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserRepositoryContext) Login(email string) (*model.Users, error) {
	user := &model.Users{}
	if err := c.cfg.Slave().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserRepositoryContext) CreateUser(data *users.CreateUserRequest) (*model.Users, error) {
	user := &model.Users{}
	user.Email = data.Email
	user.Alamat = data.Alamat
	user.Ktp = data.Ktp
	user.Nama = data.Nama
	user.NoHp = data.NoHp
	user.Password = data.Password

	if err := c.cfg.Slave().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
