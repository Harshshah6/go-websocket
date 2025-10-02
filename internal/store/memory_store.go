package store

import "sync"

type MemoryStore struct {
	mu     sync.Mutex
	Users  map[string]bool
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{Users: make(map[string]bool)}
}

func (s *MemoryStore) AddUser(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Users[name] = true
}

func (s *MemoryStore) RemoveUser(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Users, name)
}

func (s *MemoryStore) ListUsers() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	users := []string{}
	for name := range s.Users {
		users = append(users, name)
	}
	return users
}
