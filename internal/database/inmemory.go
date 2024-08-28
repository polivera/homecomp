package database

import (
	"context"
	"errors"
	"sync"
)

type InMemoryDB interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Remove(ctx context.Context, key string) error
	HasKey(ctx context.Context, key string) bool
}

type inMemoryDatabase struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewInMemoryDB() InMemoryDB {
	return &inMemoryDatabase{
		store: make(map[string]string),
	}
}

func (md *inMemoryDatabase) Set(ctx context.Context, key string, value string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		md.mu.Lock()
		defer md.mu.Unlock()
		md.store[key] = value
		return nil
	}
}

func (md *inMemoryDatabase) Get(ctx context.Context, key string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		md.mu.RLock()
		defer md.mu.Unlock()
		value, exist := md.store[key]
		if !exist {
			return "", errors.New("key does not exist")
		}
		return value, nil
	}
}

func (md *inMemoryDatabase) Remove(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		md.mu.Lock()
		defer md.mu.Unlock()
		_, exist := md.store[key]
		if !exist {
			return errors.New("key does not exist")
		}
		delete(md.store, key)
		return nil
	}
}

func (md *inMemoryDatabase) HasKey(ctx context.Context, key string) bool {
	_, err := md.Get(ctx, key)
	return err != nil
}
