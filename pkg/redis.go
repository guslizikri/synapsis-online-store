package pkg

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(ctx context.Context, host string, password string) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})

	if err = client.Ping(ctx).Err(); err != nil {
		return
	}

	return
}
