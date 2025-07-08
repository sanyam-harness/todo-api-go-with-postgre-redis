package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoService struct {
	db *pgxpool.Pool
}

func NewTodoService(db *pgxpool.Pool) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) ListTodos() ([]Todo, error) {
	rows, err := s.db.Query(context.Background(), "SELECT id, title, completed, deleted, created_at, updated_at FROM todos WHERE deleted = false")
	if err != nil {
		return nil, err
	}

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.Deleted, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(todo *Todo) *Todo {
	err := DB.QueryRow(context.Background(), `
		INSERT INTO todos (title, completed)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`, todo.Title, todo.Completed).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		return nil
	}
	return todo
}

func (s *TodoService) GetTodo(id int) (*Todo, error) {
	var todo Todo
	err := DB.QueryRow(context.Background(), `
		SELECT id, title, completed, created_at, updated_at 
		FROM todos WHERE id=$1 AND deleted=false
	`, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}

func (s *TodoService) UpdateTodo(id int, updated *Todo) (*Todo, error) {
	err := DB.QueryRow(context.Background(), `
		UPDATE todos
		SET title=$1, completed=$2, updated_at=NOW()
		WHERE id=$3 AND deleted=false
		RETURNING id, title, completed, created_at, updated_at
	`, updated.Title, updated.Completed, id).Scan(
		&updated.ID, &updated.Title, &updated.Completed, &updated.CreatedAt, &updated.UpdatedAt,
	)

	if err != nil {
		return nil, errors.New("todo not found")
	}
	return updated, nil
}

func (s *TodoService) DeleteTodo(id int) error {
	cmd, err := DB.Exec(context.Background(), `
		UPDATE todos SET deleted=true, updated_at=NOW() WHERE id=$1 AND deleted=false
	`, id)
	if err != nil || cmd.RowsAffected() == 0 {
		return errors.New("todo not found")
	}
	return nil
}
