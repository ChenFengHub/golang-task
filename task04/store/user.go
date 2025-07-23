package store

import (
	"errors"

	"github.com/ChenfengHub/golang-task/task04/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStore interface {
	RegisterUser(user *entity.User) error
	Login(user *entity.User) error
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db: db}
}

func (us *userStore) RegisterUser(user *entity.User) error {
	searchUser := entity.User{}
	us.db.Where("username = ?", user.Username).First(&searchUser)
	if searchUser.ID != 0 {
		return errors.New("用户已经存在，不可重新注册")
	}
	if err := us.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (us *userStore) Login(user *entity.User) error {
	existUser := entity.User{}
	us.db.Where("username = ?", user.Username).First(&existUser)
	if existUser.ID == 0 {
		return errors.New("账户/密码不正确！")
	}
	// 每次盐值不同，生成加密串不同，需要根据该方法判断
	err := bcrypt.CompareHashAndPassword(
		[]byte(existUser.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return errors.New("账户/密码不正确！")
	}
	user.ID = existUser.ID
	return nil
}
