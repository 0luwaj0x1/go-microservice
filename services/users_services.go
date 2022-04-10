package services

import (
	"example.com/microservices/domain"
)

var (
	UserService userService
)

type userService struct {}

func (u *userService) GetUser(userID int64) (*domain.User, error) {
	return domain.UserDao.GetUser(userID)
}