package entity

import (
	"FliQt/internals/app/enum"
	"time"
)

type Employee struct {
	ID        uint64                `gorm:"primaryKey"`
	Name      string                `gorm:"size:64;"`
	Position  enum.EmployeePosition `gorm:"default:1;"`
	Email     string                `gorm:"uniqueIndex;size:128"`
	Salary    int                   `gorm:"default:0"`
	CreatedAt time.Time             `gorm:"index"`
	UpdatedAt time.Time             `gorm:"index"`
}
