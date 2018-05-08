package memory

import "sync"

type Storage struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	value, ok := s.m[key]
	s.mu.RUnlock()
	return value, ok
}

func (s *Storage) Set(key string, value interface{}) {
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

func (s *Storage) Del(key string) {
	s.mu.Lock()
	delete(s.m, key)
	s.mu.Unlock()
}

func NewStorage() *Storage {
	return &Storage{
		m: make(map[string]interface{}),
	}
}
