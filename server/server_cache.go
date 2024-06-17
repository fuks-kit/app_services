package server

import (
	"sync"
	"time"
)

// A generic cache that stores a value of type T
//
// The Google Sheets API has a rate limit, to avoid hitting the rate limit we cache the data for 5 minutes.
type cache[T any] struct {
	mutex sync.RWMutex
	data  *T
	time  time.Time
}

// Create a new cache instance
func newCache[T any]() *cache[T] {
	return &cache[T]{}
}

// Get the cached data and whether it is valid
func (c *cache[T]) get() (*T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cached data is valid
	validCache := c.data != nil && time.Since(c.time) < 5*time.Minute

	// Return the cached data and whether it is valid
	return c.data, validCache
}

// Set the cached data
func (c *cache[T]) set(data *T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Update the cached data and the timestamp
	c.data = data
	c.time = time.Now()
}
