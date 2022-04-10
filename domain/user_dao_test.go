package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "a user with id 0 is not expected")
	assert.NotNil(t, err, "an error is expected when user id is 0")
	assert.EqualValues(t, err.Error(), "no user found")
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Josh", user.FirstName)
	assert.EqualValues(t, "Mann", user.LastName)
	assert.EqualValues(t, "test@email.com", user.Email)
}



