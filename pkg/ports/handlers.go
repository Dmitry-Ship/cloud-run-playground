package ports

import (
	"cloud-run-playground/pkg/application/usersSearch"
	"cloud-run-playground/pkg/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleRequests(usersService usersSearch.UserService) {
	http.HandleFunc("/api/users/searchByName", SearchUsers(usersService))
	http.HandleFunc("/api/users/searchById", GetUserById(usersService))
}

func SearchUsers(userService usersSearch.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Query().Get("name")

		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := userService.SearchByName(0, name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := struct {
			Users []domain.User `json:"users"`
		}{
			Users: result,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func GetUserById(userService usersSearch.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.Atoi(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := userService.GetById(idInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var response = struct {
			User domain.User `json:"user"`
		}{
			User: result,
		}

		json.NewEncoder(w).Encode(response)
	}
}
