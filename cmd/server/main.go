package main

import (
	"cloud-run-playground/pkg/adapters"
	"cloud-run-playground/pkg/application/usersSearch"
	"cloud-run-playground/pkg/ports"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := adapters.GetDatabaseConnection()

	usersRepository := adapters.NewUsersRepository(db)
	usersService := usersSearch.NewService(usersRepository)
	ports.HandleRequests(usersService)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
