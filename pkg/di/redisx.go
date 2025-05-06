package di

import (
	"FliQt/internals/app/config"
	"FliQt/pkg/redisx"
	"github.com/google/wire"
)

var WireRedisSet = wire.NewSet(
	NewRedisBundle,
	NewRedisLock,
)

func NewRedisBundle(cfg *config.Config) (*redisx.Bundle, error) {
	return redisx.NewRedisBundle(redisx.Config{
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
