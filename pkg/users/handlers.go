package users

import (
	"encoding/json"
	"net/http"
)

func HandleRequests(usersService Service) {
	http.HandleFunc("/api/users", GetUsers(usersService))
}

func GetUsers(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		result, err := userService.GetAllUsers(50)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)
	}
}
