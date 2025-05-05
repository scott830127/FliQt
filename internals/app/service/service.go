package service

import "github.com/google/wire"

var WireSet = wire.NewSet(
	wire.Struct(new(Srv), "*"),
	NewLeaveService,
)

type Srv struct {
	LeaveService ILeaveService
}
