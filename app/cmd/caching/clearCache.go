package caching

import (
	"github.com/DevAthhh/caching-proxy/initializers"
)

func ClearAllCache() error {
	if err := initializers.Client.FlushDB().Err(); err != nil {
		return err
	}

	if err := initializers.Client.FlushAll().Err(); err != nil {
		return err
	}
	return nil
}
