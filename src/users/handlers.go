package users

import (
	"Docker-Test/common"
	"net/http"
)

func HandleRequests(usersService Service) {
	http.HandleFunc("/api/users", GetUsers(usersService))
}

func GetUsers(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := userService.GetAllUsers(50)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		common.SendJSONresponse(result, w)
	}
}
