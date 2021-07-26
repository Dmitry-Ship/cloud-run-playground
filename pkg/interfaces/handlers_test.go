package interfaces

import (
	"cloud-run-playground/pkg/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockService struct{}

func (mr *MockService) SearchByName(limit int, name string) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (mr *MockService) GetById(id int) (domain.User, error) {
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

var mockService = &MockService{}

func TestSearchUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/searchByName?name=john", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchUsers(mockService))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, `{"users":[]}`, strings.TrimSpace(rr.Body.String()), "handler returned unexpected body")
}

func TestGetUserById(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/searchById?id=1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserById(mockService))

	handler.ServeHTTP(rr, req)

	user := struct {
		User domain.User `json:"user"`
	}{
		User: domain.User{
			Id:          1,
			FirstName:   "John",
			LastName:    "Doe",
			Email:       "john@doe",
			Username:    "john",
			PhoneNumber: "+1-555-555-5555",
			IPAddress:   "127.0.0.1",
			Gender:      "male",
		},
	}

	b, _ := json.Marshal(user)
	expected := string(b)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, expected, strings.TrimSpace(rr.Body.String()), "handler returned unexpected body")
}
