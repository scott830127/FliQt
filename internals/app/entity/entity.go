package entity

import (
	"FliQt/internals/app/enum"
	"FliQt/pkg/util"
	"gorm.io/gorm"
	"time"
)

func AutoMigrate() []any {
	return []any{
		new(Employee),
		new(LeaveRecord),
		new(LeaveQuota),
		new(LeaveType),
	}
}

func SeedInitialData(db *gorm.DB) error {
	var count int64

	// 員工
	db.Model(&Employee{}).Count(&count)
	if count == 0 {
		employees := []Employee{
			{ID: util.GetUniqueID(), Name: "Scott", Position: enum.EmployeePositionJr, Email: "scott@fliqt.com", Salary: 100, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: util.GetUniqueID(), Name: "Rayyy", Position: enum.EmployeePositionSr, Email: "rayyy@fliqt.com", Salary: 200, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		if err := db.Create(&employees).Error; err != nil {
			return err
		}
	}

	// 假別類型
	db.Model(&LeaveType{}).Count(&count)
	if count == 0 {
		leaveTypes := []LeaveType{
			{ID: 1, Name: "特休", Description: "年假", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Name: "病假", Description: "生病請假", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		if err := db.Create(&leaveTypes).Error; err != nil {
			return err
		}
	}

	return nil
}
