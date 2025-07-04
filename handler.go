package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *TodoService
}

func NewHandler(service *TodoService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.ListTodos())
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created := h.service.CreateTodo(&todo)
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo, err := h.service.GetTodo(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updated Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo, err := h.service.UpdateTodo(id, &updated)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.service.DeleteTodo(id); err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
