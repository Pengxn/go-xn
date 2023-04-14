package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

// Add adds a key/value pair to the cache only if the key doesn't already exist,
// or if key has expired. Returns an error otherwise.
func Add(key string, value interface{}) error {
	return c.Add(key, value, cache.DefaultExpiration)
}

// Set sets a key/value pair in the cache. If the key already exists, it will be
// overwritten.
func Set(key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}

// Get gets a value by key from the cache. It returns the value and a boolean
// indicating whether the key was found in the cache.
func Get(key string) (interface{}, bool) {
	return c.Get(key)
}

// Delete deletes a key/value pair by key from the cache whether it exists or not.
func Delete(key string) {
	c.Delete(key)
}
