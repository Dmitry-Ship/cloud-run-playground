package usersSearch

import (
	"cloud-run-playground/pkg/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct{}

func (mr *MockRepo) GetUsersByName(limit int, name string) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (mr *MockRepo) CreateUser(user domain.User) (domain.User, error) {
	return user, nil
}

func (mr *MockRepo) Find(id int) (domain.User, error) {
	return domain.User{
		Id:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@doe",
		Username:    "john",
		PhoneNumber: "+1-555-555-5555",
		IPAddress:   "127.0.0.1",
		Gender:      "male",
	}, nil
}

var mockedRepository = &MockRepo{}
var usersService = NewService(mockedRepository)

func TestGetAllUsersService(t *testing.T) {
	result, err := usersService.SearchByName(0, "")

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, 0, len(result), "The two users should be the same.")
}

func TestGetUserByIdService(t *testing.T) {
	result, err := usersService.GetById(0)

	expected := domain.User{
		Id:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@doe",
		Username:    "john",
		PhoneNumber: "+1-555-555-5555",
		IPAddress:   "127.0.0.1",
		Gender:      "male",
	}

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, expected, result, "The two users should be the same.")
}
