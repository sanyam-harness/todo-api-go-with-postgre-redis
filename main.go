package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize PostgreSQL DB
	InitDB()

	// Initialize Service Layer with DB
	service := NewTodoService(DB)

	// Initialize Redis Client for caching
	cache := NewRedisClient()

	// Initialize HTTP Handler with service and cache
	handler := NewHandler(service, cache)

	// Set up HTTP Router
	r := mux.NewRouter()
	r.HandleFunc("/todos", handler.ListTodos).Methods("GET")                 // List all todos (with Redis cache)
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")               // Create a new todo
	r.HandleFunc("/todos/{id:[0-9]+}", handler.GetTodo).Methods("GET")       // Get todo by ID
	r.HandleFunc("/todos/{id:[0-9]+}", handler.UpdateTodo).Methods("PUT")    // Update todo by ID
	r.HandleFunc("/todos/{id:[0-9]+}", handler.DeleteTodo).Methods("DELETE") // Soft delete todo by ID

	log.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
