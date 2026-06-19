package repository

import (
	"growthos/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (u *UserRepository) CreateUser(user model.User) error {
	return u.DB.Create(&user).Error
}

func (u *UserRepository) FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := u.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u *UserRepository) FindUserById(id uint) (model.User, error) {
	var user model.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *UserRepository) UpdateUserInfo(user model.User) error {
	return u.DB.Save(&user).Error
}
