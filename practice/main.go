package main

import (
	// "golang-beginner-21/handlers"
	"golang-beginner-21/practice/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /login", handlers.LoginHandler)
	http.HandleFunc("GET /user_detail", handlers.GetUserByIdHandle)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
