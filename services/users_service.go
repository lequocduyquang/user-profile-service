package services

import (
	"errors"

	"github.com/lequocduyquang/user-profile-service/domain/users"
)

var (
	// UserService exported
	UserService UserServiceInterface = &userService{}
	dao                              = &users.User{}
)

// UserServiceInterface interface
type UserServiceInterface interface {
	GetAll() (*[]users.User, error)
	GetByID(int64) (*users.User, error)
	Create(users.User) (*users.User, error)
}

type userService struct{}

func (u *userService) GetAll() (*[]users.User, error) {
	listUsers, err := dao.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return listUsers, nil
}

func (u *userService) GetByID(id int64) (*users.User, error) {
	foundUser, err := dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u *userService) Create(user users.User) (*users.User, error) {
	if err := user.Validate("register"); err != nil {
		return nil, err
	}
	existedUser, _ := user.GetUserByEmail()
	if existedUser != nil {
		return existedUser, errors.New("account with email is already exists")
	}
	user.Prepare()
	createdUser, err := user.SaveUser()
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
