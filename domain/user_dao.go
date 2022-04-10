package domain

import (
	"errors"
)

type  UserSerciveInterface interface {
	GetUser(userId int64) (*User, error)
}

var (
	UserDao UserSerciveInterface
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Josh", LastName: "Mann", Email: "test@email.com"},
	}
)

func init() {
	UserDao = &userDao{}
}

type userDao struct {}

func(u *userDao) GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New("no user found")
}