package entity

import "time"

type LeaveQuota struct {
	ID          uint64    `gorm:"primaryKey"`
	EmployeeID  uint64    `gorm:"index"`
	LeaveTypeID uint      `gorm:"index"`
	TotalHours  int       `gorm:"default:0"`
	UsedHours   int       `gorm:"default:0"`
	UpdatedAt   time.Time `gorm:"index"`

	// 預載關聯使用（不影響 DB 結構）
	Employee  Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	LeaveType LeaveType `gorm:"foreignKey:LeaveTypeID" json:"leaveType,omitempty"`
}
