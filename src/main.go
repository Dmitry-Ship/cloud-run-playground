package main

import (
	"cloud-run-playground/infrastructure"
	"cloud-run-playground/users"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := infrastructure.GetDatabaseConnection()

	usersRepository := users.NewUsersRepository(db)
	usersService := users.NewService(usersRepository)
	users.HandleRequests(usersService)

	port := os.Getenv("PORT")

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
