package repository

import "fmt"

var _ ILeaveRepository = (*LeaveRepository)(nil)

type ILeaveRepository interface {
	InsertLeave(employeeID uint, reason string, days int) error
}

type LeaveRepository struct{}

func NewLeaveRepository() ILeaveRepository {
	return &LeaveRepository{}
}

func (r *LeaveRepository) InsertLeave(employeeID uint, reason string, days int) error {
	fmt.Printf("[mock] saving leave: emp=%d reason=%s days=%d\n", employeeID, reason, days)
	return nil
}
