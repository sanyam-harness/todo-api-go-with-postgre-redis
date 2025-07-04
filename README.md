# TODO API in Go with PostgreSQL

A simple TODO REST API written in Go using PostgreSQL as a database.

---

## 📦 API Endpoints

| Method | Endpoint        | Description         |
|--------|------------------|---------------------|
| GET    | /todos           | List all todos      |
| POST   | /todos           | Create a new todo   |
| GET    | /todos/{id}      | Get todo by ID      |
| PUT    | /todos/{id}      | Update a todo       |
| DELETE | /todos/{id}      | Soft delete a todo  |

---

## 🛢️ Database Setup

Create a PostgreSQL database and table:

```sql
CREATE DATABASE tododb;
\c tododb

CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  completed BOOLEAN DEFAULT false,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
