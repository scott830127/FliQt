package di

import (
	"FliQt/internals/app/config"
	"FliQt/pkg/redisx"
	"github.com/google/wire"
)

var WireRedisSet = wire.NewSet(
	NewRedisBundle,
)

func NewRedisBundle() (*redisx.Bundle, error) {
	return redisx.NewRedisBundle(redisx.Config{
		Addr:     config.C.Redis.Addr,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,
	})
}
