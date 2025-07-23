package service

import (
	"github.com/ChenfengHub/golang-task/task04/entity"
	"github.com/ChenfengHub/golang-task/task04/store"
)

type UserService struct {
	userStore store.UserStore
}

func NewUserService(store store.UserStore) *UserService {
	return &UserService{userStore: store}
}

func (us *UserService) RegisterUser(user *entity.User) error {
	// 1. 其他验证
	// 2. 入库
	if err := us.userStore.RegisterUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) Login(user *entity.User) error {
	return us.userStore.Login(user)
}
