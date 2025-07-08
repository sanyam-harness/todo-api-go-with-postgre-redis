Great! Here's a **professional and clean `README.md`** for your repository [`todo-api-go-with-postgre-redis`](https://github.com/sanyam-harness/todo-api-go-with-postgre-redis):

---

````markdown
# TODO API in Go with PostgreSQL and Redis

A production-ready TODO REST API built using **Go**, **PostgreSQL** for persistent storage, and **Redis** for high-performance caching during list operations.  
This project follows a clean architecture with clearly separated layers: handler, service, database, and caching.

---

## 📌 Features

- ✅ Create, Read, Update, and Soft Delete TODOs
- 💾 Data stored in PostgreSQL
- 🚀 Redis cache integration for the `/todos` list endpoint
- 📦 Gorilla Mux for routing
- 🧹 Clean architecture with service abstraction

---

## 🚀 API Endpoints

| Method | Endpoint        | Description             |
|--------|------------------|-------------------------|
| GET    | `/todos`         | List all active todos (cached with Redis) |
| POST   | `/todos`         | Create a new todo       |
| GET    | `/todos/{id}`    | Get todo by ID          |
| PUT    | `/todos/{id}`    | Update a todo by ID     |
| DELETE | `/todos/{id}`    | Soft delete a todo by ID |

---

## 🧱 Database Setup (PostgreSQL)

Make sure PostgreSQL is running and execute:

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
````

---

## 🔌 Redis Setup

Ensure Redis server is running locally on the default port `6379`.
If you don't have Redis installed, you can install it using:

```bash
brew install redis
brew services start redis
```

---

## 🛠️ How to Run

1. **Clone the repository**:

```bash
git clone https://github.com/sanyam-harness/todo-api-go-with-postgre-redis.git
cd todo-api-go-with-postgre-redis
```

2. **Update the PostgreSQL DSN in `db.go` if needed**:

```go
dsn := "postgres://postgres:<your_password>@localhost:5432/tododb"
```

3. **Install dependencies**:

```bash
go mod tidy
```

4. **Run the application**:

```bash
go run main.go handler.go service.go db.go todo.go cache.go
```

> The server will start at: [http://localhost:8080](http://localhost:8080)

---

## 🧪 Test Using `curl`

### ✅ Create a TODO

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Redis Caching"}'
```

### 📋 List All TODOs (will use Redis cache after first hit)

```bash
curl http://localhost:8080/todos
```

### 🔍 Get TODO by ID

```bash
curl http://localhost:8080/todos/1
```

### ✏️ Update TODO

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Redis in Go", "completed": true}'
```

### ❌ Soft Delete TODO

```bash
curl -X DELETE http://localhost:8080/todos/1
```

---

## 🧾 Project Structure

```
├── main.go          # Starts the server and routes
├── handler.go       # Handles HTTP requests
├── service.go       # Contains business logic
├── db.go            # Connects to PostgreSQL
├── cache.go         # Connects and manages Redis cache
├── todo.go          # Defines the Todo model
```

---

## 📚 Tech Stack

* **Go (Golang)**
* **PostgreSQL**
* **Redis**
* **Gorilla Mux**

