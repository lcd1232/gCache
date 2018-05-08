package memory

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExpired  = errors.New("key expired")
)

type Item struct {
	Value      interface{}
	Expiration int64
}

type TTLStorage struct {
	m  map[string]Item
	mu sync.RWMutex
	de time.Duration
	ci time.Duration
}

func (s *TTLStorage) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	item, ok := s.m[key]
	if !ok {
		s.mu.RUnlock()
		return nil, false
	}
	if time.Now().UnixNano() > item.Expiration {
		s.mu.RUnlock()
		return nil, false
	}
	s.mu.RUnlock()
	return item.Value, ok
}

func (s *TTLStorage) Set(key string, value interface{}) {
	s.mu.Lock()
	e := time.Now().Add(s.de).UnixNano()
	item := Item{
		Value:      value,
		Expiration: e,
	}
	s.m[key] = item
	s.mu.Unlock()
}

func (s *TTLStorage) Del(key string) {
	s.mu.Lock()
	delete(s.m, key)
	s.mu.Unlock()
}

func (s *TTLStorage) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	s.mu.Lock()
	item := Item{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
	s.m[key] = item
	s.mu.Unlock()
}

func (s *TTLStorage) SetTTL(key string, ttl time.Duration) error {
	s.mu.Lock()
	item, ok := s.m[key]
	if !ok {
		s.mu.Unlock()
		return ErrKeyNotFound
	}
	if time.Now().UnixNano() > item.Expiration {
		return ErrKeyExpired
	}
	item.Expiration = time.Now().Add(ttl).UnixNano()
	s.m[key] = item
	s.mu.Unlock()
	return nil
}

func (s *TTLStorage) GetWithTTL(key string) (interface{}, time.Time, bool) {
	s.mu.RLock()
	item, ok := s.m[key]
	if !ok {
		s.mu.RUnlock()
		return nil, time.Time{}, false
	}
	if time.Now().UnixNano() > item.Expiration {
		s.mu.RUnlock()
		return nil, time.Time{}, false
	}
	s.mu.RUnlock()
	return item.Value, time.Unix(0, item.Expiration), true
}

func NewTTLStorage(defaultExpiration, cleanupInterval time.Duration) *TTLStorage {
	return &TTLStorage{
		m:  make(map[string]Item),
		de: defaultExpiration,
		ci: cleanupInterval,
	}
}
