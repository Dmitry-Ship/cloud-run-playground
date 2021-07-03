package main

import (
	"cloud-run-playground/pkg/storage"
	"cloud-run-playground/pkg/users"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := storage.GetDatabaseConnection()

	usersRepository := users.NewUsersRepository(db)
	usersService := users.NewService(usersRepository)
	users.HandleRequests(usersService)

	port := os.Getenv("PORT")

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
