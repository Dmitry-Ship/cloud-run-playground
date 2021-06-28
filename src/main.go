package main

import (
	"Docker-Test/infrastructure"
	"Docker-Test/users"
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

	http.Handle("/", http.FileServer(http.Dir("./static")))

	port := os.Getenv("PORT")
	fmt.Println("PORT " + port)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
