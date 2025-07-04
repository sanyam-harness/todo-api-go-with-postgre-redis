package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dsn := "postgres://postgres:admin123@localhost:5432/tododb"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	log.Println("Connected to PostgreSQL database")
}
