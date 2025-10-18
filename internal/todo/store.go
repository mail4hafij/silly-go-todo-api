package todo

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var ErrNotFound = errors.New("todo not found")

type Store interface {
	Create(ctx context.Context, t Todo) (Todo, error)
	Get(ctx context.Context, id int) (Todo, error)
	Update(ctx context.Context, t Todo) (Todo, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]Todo, error)
}

type InMemoryStore struct {
	mu     sync.RWMutex
	nextID int
	items  map[int]Todo
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		nextID: 1,
		items:  make(map[int]Todo),
	}
}

func (s *InMemoryStore) Create(ctx context.Context, t Todo) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	t.ID = s.nextID
	s.nextID++
	t.CreatedAt = now
	t.UpdatedAt = now
	s.items[t.ID] = t
	return t, nil
}

func (s *InMemoryStore) Get(ctx context.Context, id int) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.items[id]
	if !ok {
		return Todo{}, ErrNotFound
	}
	return t, nil
}

func (s *InMemoryStore) Update(ctx context.Context, t Todo) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	existing, ok := s.items[t.ID]
	if !ok {
		return Todo{}, ErrNotFound
	}
	t.CreatedAt = existing.CreatedAt
	t.UpdatedAt = time.Now().UTC()
	s.items[t.ID] = t
	return t, nil
}

func (s *InMemoryStore) Delete(ctx context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return ErrNotFound
	}
	delete(s.items, id)
	return nil
}

func (s *InMemoryStore) List(ctx context.Context) ([]Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Todo, 0, len(s.items))
	for _, t := range s.items {
		out = append(out, t)
	}
	return out, nil
}
