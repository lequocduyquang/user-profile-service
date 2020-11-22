package services

import "github.com/lequocduyquang/user-profile-service/domain/users"

var (
	// UserService exported
	UserService UserServiceInterface = &userService{}
)

// UserServiceInterface interface
type UserServiceInterface interface {
	GetAll() (*[]users.User, error)
}

type userService struct{}

func (u *userService) GetAll() (*[]users.User, error) {
	dao := &users.User{}
	listUsers, err := dao.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return listUsers, nil
}
