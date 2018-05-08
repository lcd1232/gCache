package storage

import "time"

type BaseStorage interface {
	Get(key string) (ok bool, value interface{})
	Set(key string, value interface{})
	Del(key string) error
}

type TTLStorage interface {
	BaseStorage
	SetWithTTL(key string, value interface{}, ttl time.Duration)
	SetTTL(key string, ttl time.Duration) error
}

type Storage interface {
	BaseStorage
}