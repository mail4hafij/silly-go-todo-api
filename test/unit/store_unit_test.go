package unit

import (
	"context"
	"testing"

	"todo-api/internal/todo"
)

func TestInMemoryStore_BasicOps(t *testing.T) {
	s := todo.NewInMemoryStore() // exported constructor
	ctx := context.Background()

	// Create
	created, err := s.Create(ctx, todo.Todo{Title: "unit test"})
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if created.ID == 0 {
		t.Fatalf("expected id assigned")
	}

	// Get
	got, err := s.Get(ctx, created.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if got.Title != "unit test" {
		t.Fatalf("unexpected title: %q", got.Title)
	}

	// Update
	got.Title = "updated"
	updated, err := s.Update(ctx, got)
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}
	if updated.Title != "updated" {
		t.Fatalf("update didn't persist")
	}

	// List
	list, err := s.List(ctx)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected list length 1, got %d", len(list))
	}

	// Delete
	if err := s.Delete(ctx, created.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	_, err = s.Get(ctx, created.ID)
	if err == nil {
		t.Fatalf("expected not found after delete")
	}
}
