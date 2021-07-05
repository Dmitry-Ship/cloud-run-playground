package usersSearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct{}

func (mr *MockRepo) GetUsersByName(limit int, name string) ([]User, error) {
	return []User{}, nil
}

func (mr *MockRepo) CreateUser(user User) (User, error) {
	return user, nil
}

func (mr *MockRepo) GetUserById(id int) (User, error) {
	return User{}, nil
}

var mockedRepository = &MockRepo{}
var userService = NewService(mockedRepository)

func TestGetAllUsersService(t *testing.T) {
	result, err := userService.SearchByName(0, "")

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, 0, len(result), "The two users should be the same.")
}
