package redisx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

type Bundle struct {
	Client *redis.Client
	Locker *Locker
}

func NewRedisBundle(cfg Config) (*Bundle, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}
	return &Bundle{Client: client}, nil
}
