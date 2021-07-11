package main

import (
	"cloud-run-playground/pkg/domain/usersSearch"
	"cloud-run-playground/pkg/rest"
	"cloud-run-playground/pkg/storage"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := storage.GetDatabaseConnection()

	usersRepository := storage.NewUsersRepository(db)
	usersService := usersSearch.NewService(usersRepository)
	rest.HandleRequests(usersService)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
