package dto

import (
	"FliQt/internals/app/enum"
	"time"
)

type LeaveRecordCreateCommand struct {
	ID          uint64                 `json:"-"`
	EmployeeID  uint64                 `json:"employeeID"`
	LeaveTypeID uint                   `json:"leaveTypeID"`
	StartTime   time.Time              `json:"startTime"`
	EndTime     time.Time              `json:"endTime"`
	Hours       int                    `json:"hours"`
	Reason      string                 `json:"reason"`
	Status      enum.LeaveRecordStatus `json:"-"`
}

type LeaveTypeResult struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}
