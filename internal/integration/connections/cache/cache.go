package connection

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/dock-tech/notes-api/internal/config/properties"
	"github.com/redis/go-redis/v9"
)

func cache(addr string) *redis.Client {

	opts := &redis.Options{
		Addr:       addr,
		MaxRetries: properties.GetCacheMaxRetries(),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
		DialTimeout: properties.GetCacheTimeout(),
	}

	if properties.GetEnv() == "local" {
		opts.TLSConfig = nil
	}

	client := redis.NewClient(opts)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}

	return client
}

func CacheGet() *redis.Client {
	return cache(properties.GetCacheGetHost())
}

func CacheSet() *redis.Client {
	return cache(properties.GetCacheSetHost())
}

func DisconnectCache(clients ...*redis.Client) error {
	for _, client := range clients {
		err := client.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
