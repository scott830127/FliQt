package entity

import (
	"FliQt/internals/app/enum"
	"time"
)

type LeaveRecord struct {
	ID          uint64                 `gorm:"primaryKey"`
	EmployeeID  uint64                 `gorm:"index"`
	LeaveTypeID uint                   `gorm:"index"`
	StartTime   time.Time              `gorm:"index"`
	EndTime     time.Time              `gorm:"index"`
	Hours       int                    `gorm:"default:0"`
	Reason      string                 `gorm:"size:128;"`
	Status      enum.LeaveRecordStatus `gorm:"default:1;"`
	CreatedAt   time.Time              `gorm:"index"`
	UpdatedAt   time.Time              `gorm:"index"`

	// 預載關聯使用
	Employee  Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	LeaveType LeaveType `gorm:"foreignKey:LeaveTypeID" json:"leaveType,omitempty"`
}
