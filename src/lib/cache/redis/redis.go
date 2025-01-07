package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Pengxn/go-xn/src/config"
)

const DefaultExpiration = 5 * time.Minute

var (
	// UseRedis indicates whether to use redis.
	UseRedis bool

	r *redis.Client
)

func init() {
	if config.Config.Redis.URL != "" {
		UseRedis = true
	}

	r = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.URL,
		Username: config.Config.Redis.Username,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
}

// Add adds a key/value pair to the redis only if the key doesn't already exist,
// or if the key has expired. Returns an error otherwise.
func Add(key string, value any) error {
	return r.Set(context.Background(), key, value, DefaultExpiration).Err()
}

// Set sets a key/value pair in the reids. If the key already exists, it will be
// overwritten.
func Set(key string, value any) {
	r.Set(context.Background(), key, value, DefaultExpiration)
}

// SetWithExpiration sets a key/value pair in the reids with a given expiration.
// If the key already exists, it will be overwritten. If expiration is 0, the
// default expiration time will be used.
func SetWithExpiration(key string, value any, expiration time.Duration) {
	r.Set(context.Background(), key, value, expiration)
}

// Get gets a value by key from the redis. It returns the value and a boolean
// indicating whether the key was found in the redis.
func Get(key string) (any, bool) {
	result, err := r.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		return "", false
	}

	return result, true
}

// Delete deletes a key/value pair by key from the redis whether it exists or not.
func Delete(key string) {
	r.Del(context.Background(), key)
}
