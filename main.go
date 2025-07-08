package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// ✅ Initialize PostgreSQL (DB is a global variable from db.go)
	InitDB()

	// ✅ Initialize Redis (returns *redis.Client)
	rdb := InitRedis()

	// ✅ Check Redis connection
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}
	log.Printf("✅ Redis ping response: %s", pong)

	// ✅ Create the service with both PostgreSQL and Redis clients
	service := NewTodoService(DB, rdb)

	// ✅ Create handler with the service injected
	handler := NewHandler(service)

	// ✅ Set up Gorilla Mux routes
	r := mux.NewRouter()
	r.HandleFunc("/todos", handler.ListTodos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.DeleteTodo).Methods("DELETE")

	// ✅ Start the HTTP server
	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
