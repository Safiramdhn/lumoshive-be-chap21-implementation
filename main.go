package main

import (
	"fmt"
	"golang-beginner-21/handlers"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	// Register routes
	serverMux.HandleFunc("/login", handlers.LoginHandler)
	serverMux.HandleFunc("/register", handlers.CreateUserHandler)

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", serverMux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
