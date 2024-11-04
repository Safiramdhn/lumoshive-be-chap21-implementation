package main

import (
	"fmt"
	"golang-beginner-21/handlers"
	"golang-beginner-21/middleware"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	// Register routes
	authMux := http.NewServeMux()
	authMux.HandleFunc("/login", handlers.LoginHandler)
	authMux.HandleFunc("/register", handlers.CreateUserHandler)

	todoMux := http.NewServeMux()
	todoMux.HandleFunc("/getall", handlers.GetTodosHandler)
	todoMux.HandleFunc("/create-todo", handlers.CreateTodoHandler)
	todoMux.HandleFunc("/update-todo", handlers.UpdateTodoHandler)
	todoMux.HandleFunc("/delete-todo", handlers.DeleteTodoHandler)
	todoMux.HandleFunc("/get-todo-count", handlers.GetTodoCountHandler)

	middleware := middleware.Middleware(todoMux)
	serverMux.Handle("/", authMux)
	serverMux.Handle("/todo/", http.StripPrefix("/todo", middleware))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", serverMux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
