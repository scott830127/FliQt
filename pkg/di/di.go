package di

import "github.com/google/wire"

var WireSet = wire.NewSet(
	WireGormSet,
	WireRedisSet,
)
