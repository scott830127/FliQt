package di

import (
	"FliQt/internals/app/config"
	"FliQt/pkg/redisx"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var WireRedisSet = wire.NewSet(
	NewRedisLock,
	NewRedisClient,
	wire.Struct(new(redisx.Bundle), "*"),
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	return redisx.NewClient(redisx.Config{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}

func NewRedisLock(cfg *config.Config) (*redisx.Locker, func(), error) {
	return redisx.NewLocker(redisx.Config{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}
