package domain

import "errors"


var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Josh", LastName: "Mann", Email: "test@email.com"},
	}
)


func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New("no user found")
}