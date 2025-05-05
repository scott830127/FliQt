package entity

type Employee struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:64;not null"`
	Position string `gorm:"size:64"`
	Email    string `gorm:"uniqueIndex;size:128"`
	Salary   int
}
