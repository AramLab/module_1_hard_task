package cache

import (
	"sync"
)

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Cache — реализация кеша с использованием map и RWMutex
type Cache struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	return &Cache{
		store: make(map[string]string),
	}
}

// Set записывает значение по ключу
func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[k] = v
}

// Get возвращает значение по ключу
func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.store[k]
	return v, ok
}
