package services

import (
	"errors"
	"net/http"
	"testing"

	"example.com/microservices/domain"
	"github.com/stretchr/testify/assert"
)


type UsersDaoStub struct {}

var userMockGenerator func(userId int64) (*domain.User, error)

func(u *UsersDaoStub) GetUser(userId int64) (*domain.User, error) {
	return userMockGenerator(userId)
}

func init(){
	domain.UserDao = &UsersDaoStub{}
}

func TestGetUserNoUserFoundFromDB(t *testing.T){
	userMockGenerator = func(userId int64) (*domain.User, error) {
		return nil, errors.New("no user found")
	}
	user, err := UserService.GetUser(1)
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusNotFound, 404)
	assert.EqualValues(t, err.Error(), "no user found")
}

func TestGetUserNoError(t *testing.T) {
	userMockGenerator = func(userId int64) (*domain.User, error) {
		return &domain.User{
			FirstName: "Josh",
			LastName: "Mann",
			Email: "Josh@test.com",
		}, nil
	}
	user, err := UserService.GetUser(55)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, 200)
	assert.EqualValues(t, user.FirstName, "Josh")
	assert.EqualValues(t, user.LastName, "Mann")
	assert.EqualValues(t, user.Email, "Josh@test.com")
}