package repository

import "github.com/google/wire"

var WireSet = wire.NewSet(
	wire.Struct(new(Repo), "*"),
	NewLeaveRepository,
)

type Repo struct {
	LeaveRepository ILeaveRepository
}
