package users

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockService struct{}

func (mr *MockService) GetAllUsers(limit int) ([]User, error) {
	return []User{}, nil
}

var mockService = &MockService{}

func TestGetUsersHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers(mockService))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code")
	assert.Equal(t, "[]", strings.TrimSpace(rr.Body.String()), "handler returned unexpected body")
}
