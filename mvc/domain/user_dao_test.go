package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, appError := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, appError, "we were expecting an error when id is 0")
	assert.EqualValues(t, http.StatusNotFound, appError.Status, "we were expecting an error when user is not found")
	assert.EqualValues(t, "not_found", appError.Code)
	assert.EqualValues(t, "user 0 was not found", appError.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, appError := GetUser(123)

	assert.Nil(t, appError)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Gabriel", user.FirstName)
	assert.EqualValues(t, "Bussolo", user.LastName)
	assert.EqualValues(t, "contact@gabrielbussolo.dev", user.Email)
}
