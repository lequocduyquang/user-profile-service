package users

import (
	"errors"
	"strings"
)

// User define model
type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	Password string `json:"password"`
}

// Users define model
type Users []User

// Validate function
func (u *User) Validate() error {
	u.UserName = strings.TrimSpace(u.UserName)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	if u.UserName == "" {
		return errors.New("Invalid username")
	}
	if u.Email == "" {
		return errors.New("Invalid email")
	}
	if u.Password == "" {
		return errors.New("Invalid password")
	}
	return nil
}
