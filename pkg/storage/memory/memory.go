package memory

import "sync"

type MemoryStorage struct {
	m map[string]interface{}
	mu sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		m: make(map[string]interface{}),
	}
}