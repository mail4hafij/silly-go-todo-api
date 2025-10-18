# 🧩 Todo API — Go Project

A simple, clean RESTful API written in Go that manages a list of todos.  
This project demonstrates idiomatic Go structure with `cmd/`, `internal/`, and organized `test/` folders for unit and integration testing.

## 🚀 Getting Started

### 1. Install Go

Make sure Go (version 1.21 or later) is installed:

```bash
go version
```

### 2. Setup Project

Clone or unzip the project and open it in your terminal:

```bash
cd todo-api
```

Initialize dependencies:

```bash
go mod tidy
```

## ⚙️ 3. Run the API Server

To start the HTTP server:

```bash
go run ./cmd/server
```

## 📡 4. Test the API Endpoints

You can interact with the API using Postman, curl, or any REST client.

### ➕ Create a Todo

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"Buy milk"}' http://localhost:8080/todos
```

### 📋 List All Todos

```bash
curl http://localhost:8080/todos
```

### ✅ Update a Todo

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"completed":true}' http://localhost:8080/todos/1
```

### ❌ Delete a Todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## 🧪 5. Run Tests

Run all tests:

```bash
go test ./...
```

Run unit tests only:

```bash
go test ./test/unit/...
```

Run integration tests only:

```bash
go test ./test/integration/...
```

## 🧹 6. Stop the Server

Press Ctrl + C to stop the server.

## 🧱 Project Structure

```
todo-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── todo/
│       ├── store.go
│       └── handlers.go
├── test/
│   ├── unit/
│   │   └── store_unit_test.go
│   └── integration/
│       └── api_integration_test.go
├── go.mod
└── README.md
```

## 🧭 Notes

- The project uses an in-memory store (no database) for simplicity.  
- All tests depend only on Go's standard library.  
- Can be extended to use a real database if needed.  
