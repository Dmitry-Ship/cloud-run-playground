package rest

import (
	"cloud-run-playground/pkg/domain/usersSearch"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockService struct{}

func (mr *MockService) SearchByName(limit int, name string) ([]usersSearch.User, error) {
	return []usersSearch.User{}, nil
}

var mockService = &MockService{}

func TestSearchUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users?name=john", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchUsers(mockService))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, "[]", strings.TrimSpace(rr.Body.String()), "handler returned unexpected body")
}
