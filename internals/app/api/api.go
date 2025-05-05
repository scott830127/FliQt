package api

import "github.com/google/wire"

var WireSet = wire.NewSet(
	wire.Struct(new(Api), "*"),
	NewLeaveAPI,
)

type Api struct {
	LeaveAPI ILeaveAPI
}
