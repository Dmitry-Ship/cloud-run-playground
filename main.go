package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api", handleAPI)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	fmt.Println("Listening to: http://" + host + ":" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
