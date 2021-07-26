package main

import (
	"cloud-run-playground/pkg/application/usersSearch"
	"cloud-run-playground/pkg/interfaces"
	"cloud-run-playground/pkg/persistance"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := persistance.GetDatabaseConnection()

	usersRepository := persistance.NewUsersRepository(db)
	usersService := usersSearch.NewService(usersRepository)
	interfaces.HandleRequests(usersService)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
