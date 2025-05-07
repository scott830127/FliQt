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

type LeaveQuotaQuery struct {
	EmployeeID  uint64 `json:"employeeID" form:"employeeID"`
	LeaveTypeID uint   `json:"leaveTypeID" form:"leaveTypeID"`
}

type LeaveQuotaResult struct {
	EmployeeID    uint64                `json:"employeeID"`
	EmployeeName  string                `json:"employeeName"`
	Position      enum.EmployeePosition `json:"position"`
	Email         string                `json:"email"`
	Salary        int                   `json:"salary"`
	LeaveTypeID   uint                  `json:"leaveTypeID"`
	LeaveTypeName string                `json:"leaveTypeName"`
	Description   string                `json:"description"`
	TotalHours    int                   `json:"totalHours"`
	UsedHours     int                   `json:"usedHours"`
	UpdatedAt     time.Time             `json:"updatedAt"`
}

type LeaveRecordQuery struct {
	RecordID    uint64                 `json:"recordID" form:"recordID"`
	EmployeeID  uint64                 `json:"employeeID" form:"employeeID"`
	LeaveTypeID uint                   `json:"leaveTypeID" form:"leaveTypeID"`
	Hours       int                    `json:"hours" form:"hours"`
	Status      enum.LeaveRecordStatus `json:"status" form:"status"`
}

type LeaveRecordResult struct {
	RecordID      uint64                 `json:"recordID"`
	EmployeeID    uint64                 `json:"employeeID"`
	EmployeeName  string                 `json:"employeeName"`
	Position      enum.EmployeePosition  `json:"position"`
	Email         string                 `json:"email"`
	Salary        int                    `json:"salary"`
	LeaveTypeID   uint                   `json:"leaveTypeID"`
	LeaveTypeName string                 `json:"leaveTypeName"`
	Description   string                 `json:"description"`
	StartTime     time.Time              `json:"startTime"`
	EndTime       time.Time              `json:"endTime"`
	Hours         int                    `json:"hours"`
	Reason        string                 `json:"reason"`
	Status        enum.LeaveRecordStatus `json:"status"`
	CreatedAt     time.Time              `json:"createdAt"`
	UpdatedAt     time.Time              `json:"updatedAt"`
}

type LeaveRecordUpdateCommand struct {
	RecordID uint64                 `json:"recordID"`
	Status   enum.LeaveRecordStatus `json:"status"`
}

type LeaveQuotaUpdateCommand struct {
	EmployeeID  uint64 `json:"employeeID" `
	LeaveTypeID uint   `json:"leaveTypeID" `
	UsedHours   int    `json:"usedHours"`
}

type LeaveTypeResult struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}
