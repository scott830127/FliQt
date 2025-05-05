package entity

type LeaveRecord struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Reason     string
	Days       int
}
