package storage

import "time"

type BaseStorage interface {
	Get(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
	Del(key string)
}

type TTLStorage interface {
	BaseStorage
	SetWithTTL(key string, value interface{}, ttl time.Duration)
	SetTTL(key string, ttl time.Duration) error
	GetWithTTL(key string) (value interface{}, ttl time.Time, ok bool)
}

type Storage interface {
	BaseStorage
}
