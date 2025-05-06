package entity

import "time"

type LeaveType struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:64;"`
	Description string    `gorm:"size:128;"`
	IsActive    bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
}
