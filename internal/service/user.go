package service

import (
	"errors"
	"growthos/internal/model"
	"growthos/internal/repository"
	"growthos/internal/request"
	"growthos/pkg/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		Repo: repo,
	}
}

func (u *UserService) Register(user model.User) error {
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	user.Password = password
	return u.Repo.CreateUser(user)
}

func (u *UserService) Login(req request.LoginReq) (model.User, error) {
	user, err := u.Repo.FindUserByUsername(req.Username)
	if err != nil {
		return user, errors.New("获取用户信息失败")
	}
	ok := utils.CheckPassword(req.Password, user.Password)
	if !ok {
		return user, errors.New("密码错误")
	}
	return user, nil
}
