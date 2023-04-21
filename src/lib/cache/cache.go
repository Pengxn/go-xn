package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/Pengxn/go-xn/src/lib/cache/redis"
)

var c = cache.New(redis.DefaultExpiration, 10*time.Minute)

// Add adds a key/value pair to the cache only if the key doesn't already exist,
// or if the key has expired. Returns an error otherwise.
func Add(key string, value interface{}) error {
	if redis.UseRedis {
		return redis.Add(key, value)
	}

	return c.Add(key, value, cache.DefaultExpiration)
}

// Set sets a key/value pair in the cache. If the key already exists, it will be
// overwritten.
func Set(key string, value interface{}) {
	if redis.UseRedis {
		redis.Set(key, value)
		return
	}

	c.Set(key, value, cache.DefaultExpiration)
}

// SetWithExpiration sets a key/value pair in the cache with a given expiration.
// If the key already exists, it will be overwritten. If expiration is 0, the
// default expiration time will be used.
func SetWithExpiration(key string, value interface{}, expiration time.Duration) {
	if redis.UseRedis {
		redis.SetWithExpiration(key, value, expiration)
		return
	}

	c.Set(key, value, expiration)
}

// Get gets a value by key from the cache. It returns the value and a boolean
// indicating whether the key was found in the cache.
func Get(key string) (interface{}, bool) {
	if redis.UseRedis {
		return redis.Get(key)
	}

	return c.Get(key)
}

// Delete deletes a key/value pair by key from the cache whether it exists or not.
func Delete(key string) {
	if redis.UseRedis {
		redis.Delete(key)
		return
	}

	c.Delete(key)
}
