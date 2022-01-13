// package middleware houses all logic that enables HTTP middlewares like loggers, auth,
// caching, tracing etc.
package cache

import (
	"sync"
	"time"
)

// Item is a single cache item, which is a struct with content in bytes and expiration.
type Item struct {
	Content    []byte
	Expiration int64
}

// Expired checks if the item has expired.
func (i Item) Expired() bool {
	if i.Expiration == 0 {
		return false
	}

	return time.Now().UnixNano() > i.Expiration
}

// Cache is an in-memory cache that is thread-safe.
type Cache struct {
	Items map[string]Item
	mutex *sync.RWMutex
}

// NewCache creates an in-memory cache instance.
func NewCache() *Cache {
	return &Cache{
		Items: make(map[string]Item),
		mutex: &sync.RWMutex{},
	}
}

// Get fetches a value from the cache. If expired, will delete it from
// the cache and return `nil`.
func (c Cache) Get(key string) []byte {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item := c.Items[key]
	if item.Expired() {
		delete(c.Items, key)
		return nil
	}

	return item.Content
}

// Sets will set cached content by key and duration.
func (c Cache) Set(key string, content []byte, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Items[key] = Item{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}
