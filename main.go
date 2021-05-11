package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type response struct {
	Data string `json:"data"`
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := response{Data: "Hello World"}

	json.NewEncoder(w).Encode(res)
}

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api", handleAPI)

	port := os.Getenv("PORT")
	fmt.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
