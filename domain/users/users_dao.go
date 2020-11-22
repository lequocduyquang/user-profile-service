package users

import (
	"github.com/lequocduyquang/user-profile-service/db"
)

// SaveUser adds a user to the database
func (u *User) SaveUser() (*User, error) {
	var err error
	err = db.Client.Debug().Create(&u).Error
	if err != nil {
		return &User{}, nil
	}
	return u, nil
}

// GetUser returns a user based on email
func (u *User) GetUser() (*User, error) {
	account := &User{}
	if err := db.Client.Debug().Table("users").Where("email = ?", u.Email).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

// GetAllUsers returns a list of all the user
func (u *User) GetAllUsers() (*[]User, error) {
	users := []User{}
	if err := db.Client.Debug().Table("users").Find(&users).Error; err != nil {
		return &[]User{}, nil
	}
	return &users, nil
}
