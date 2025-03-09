package caching

import (
	"time"

	"github.com/DevAthhh/caching-proxy/initializers"
)

func GetCache(key string) (string, error) {
	return initializers.Client.Get(key).Result()
}

func SetCache(key, value string) error {
	if err := initializers.Client.Set(key, value, time.Hour*24).Err(); err != nil {
		return err
	}
	return nil
}
