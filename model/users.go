package model

import (
	"myapp/config"
)

type Users struct {
	Email    string `json:"email" form:"email" gorm:"primaryKey"`
	Nama     string `json:"nama" form:"nama"`
	NoHp     string `json:"no_hp" form:"no_hp"`
	Alamat   string `json:"alamat" form:"alamat"`
	Password string `json:"password" form:"password"`
	Ktp      string `json:"ktp" form:"ktp"`
}

func (user *Users) CreateUser() error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateUser(email string) error {
	if err := config.DB.Model(&Users{}).Where("email = ?", email).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByEmail(email string) (Users, error) {
	var user Users
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetAll(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Where("email LIKE ? OR nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)

	return users, result.Error
}
