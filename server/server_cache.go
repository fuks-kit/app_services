package server

import (
	"sync"
	"time"
)

type cache[T any] struct {
	mutex sync.RWMutex
	data  *T
	time  time.Time
}

func newCache[T any]() *cache[T] {
	// Returning a new instance of a generic cache
	return &cache[T]{}
}

func (c *cache[T]) get() (*T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cached data is valid
	validCache := c.data != nil && time.Since(c.time) < 5*time.Minute

	// Return the cached data and whether it is valid
	return c.data, validCache
}

func (c *cache[T]) set(data *T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Update the cached data and the timestamp
	c.data = data
	c.time = time.Now()
}
