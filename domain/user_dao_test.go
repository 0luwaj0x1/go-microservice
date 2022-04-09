package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "a user with id 0 is not expected")
	assert.NotNil(t, err, "an error is expected when user id is 0")
	assert.EqualValues(t, err.Error(), "no user found")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}



