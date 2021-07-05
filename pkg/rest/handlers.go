package rest

import (
	"cloud-run-playground/pkg/domain/usersSearch"
	"encoding/json"
	"net/http"
)

func HandleRequests(usersService usersSearch.UserService) {
	http.HandleFunc("/api/users", SearchUsers(usersService))
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

		json.NewEncoder(w).Encode(result)
	}
}
