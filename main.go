package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()

	service := NewTodoService()
	handler := NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/todos", handler.ListTodos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.DeleteTodo).Methods("DELETE")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
