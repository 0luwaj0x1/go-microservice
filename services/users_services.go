package services

import (
	"example.com/microservices/domain"
)


func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}