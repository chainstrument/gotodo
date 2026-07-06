package todo

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("todo not found")

type Store struct {
	mu    sync.RWMutex
	todos map[string]Todo
}

func NewStore() *Store {
	return &Store{todos: make(map[string]Todo)}
}

func (s *Store) List() []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]Todo, 0, len(s.todos))
	for _, t := range s.todos {
		list = append(list, t)
	}
	return list
}

func (s *Store) Get(id string) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.todos[id]
	if !ok {
		return Todo{}, ErrNotFound
	}
	return t, nil
}

func (s *Store) Create(title, description string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	t := Todo{
		ID:          newID(),
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	s.todos[t.ID] = t
	return t
}

func (s *Store) Update(id, title, description string, done bool) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.todos[id]
	if !ok {
		return Todo{}, ErrNotFound
	}
	t.Title = title
	t.Description = description
	t.Done = done
	t.UpdatedAt = time.Now()
	s.todos[id] = t
	return t, nil
}

func (s *Store) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.todos[id]; !ok {
		return ErrNotFound
	}
	delete(s.todos, id)
	return nil
}

func newID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
