package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type TodoService struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewTodoService(db *pgxpool.Pool, rdb *redis.Client) *TodoService {
	return &TodoService{
		db:  db,
		rdb: rdb,
	}
}

func (s *TodoService) ListTodos() ([]*Todo, error) {
	ctx := context.Background()
	cacheKey := "todos"

	// Try to get from Redis
	if cached, err := s.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var todos []*Todo
		if err := json.Unmarshal([]byte(cached), &todos); err == nil {
			log.Println("Fetched from Redis")
			return todos, nil
		}
	}

	// Else fetch from DB
	rows, err := s.db.Query(ctx, "SELECT id, title, completed, deleted, created_at, updated_at FROM todos WHERE deleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		t := &Todo{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.Deleted, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	// Store in Redis
	data, _ := json.Marshal(todos)
	s.rdb.Set(ctx, cacheKey, data, 10*time.Second)
	fmt.Println("Data stored in Redis")

	return todos, nil
}

func (s *TodoService) CreateTodo(todo *Todo) (*Todo, error) {
	now := time.Now()
	ctx := context.Background()
	err := s.db.QueryRow(ctx, "INSERT INTO todos (title, completed, deleted, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		todo.Title, todo.Completed, false, now, now).Scan(&todo.ID)
	if err != nil {
		return nil, err
	}
	todo.CreatedAt = now
	todo.UpdatedAt = now

	// Invalidate cache
	s.rdb.Del(context.Background(), "todos")
	log.Println("Cache invalidated")

	return todo, nil
}

func (s *TodoService) GetTodo(id int) (*Todo, error) {
	ctx := context.Background()
	todo := &Todo{}
	err := s.db.QueryRow(ctx, "SELECT id, title, completed, deleted, created_at, updated_at FROM todos WHERE id = $1 AND deleted = false", id).
		Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.Deleted, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) UpdateTodo(id int, updated *Todo) (*Todo, error) {
	now := time.Now()
	ctx := context.Background()
	res, err := s.db.Exec(ctx, "UPDATE todos SET title=$1, completed=$2, updated_at=$3 WHERE id=$4 AND deleted=false",
		updated.Title, updated.Completed, now, id)
	if err != nil {
		return nil, err
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("todo not found")
	}

	// Invalidate cache
	s.rdb.Del(context.Background(), "todos")

	return s.GetTodo(id)
}

func (s *TodoService) DeleteTodo(id int) error {
	now := time.Now()
	res, err := s.db.Exec(context.Background(), "UPDATE todos SET deleted=true, updated_at=$1 WHERE id=$2 AND deleted=false", now, id)
	if err != nil {
		return err
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("todo not found")
	}

	// Invalidate cache
	s.rdb.Del(context.Background(), "todos")

	return nil
}
