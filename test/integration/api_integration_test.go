package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-api/internal/todo"
)

func setupHandler() http.Handler {
	store := todo.NewInMemoryStore()
	h := todo.NewHandler(store)
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	return mux
}

func TestTodoAPI_CreateGetUpdateDelete(t *testing.T) {
	h := setupHandler()

	// Create
	reqBody := map[string]string{"title": "integration test"}
	b, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(b))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create failed: code=%d body=%s", rec.Code, rec.Body.String())
	}
	var created todo.Todo
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatalf("unmarshal created: %v", err)
	}

	// Get
	req = httptest.NewRequest(http.MethodGet, "/todos/"+fmt.Sprintf("%d", created.ID), nil)
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("get failed: %d", rec.Code)
	}

	// Update
	update := map[string]bool{"completed": true}
	b, _ = json.Marshal(update)
	req = httptest.NewRequest(http.MethodPatch, "/todos/"+fmt.Sprintf("%d", created.ID), bytes.NewReader(b))
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("update failed: %d", rec.Code)
	}

	// Delete
	req = httptest.NewRequest(http.MethodDelete, "/todos/"+fmt.Sprintf("%d", created.ID), nil)
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		body, _ := io.ReadAll(rec.Body)
		t.Fatalf("delete failed: %d body=%s", rec.Code, string(body))
	}
}
