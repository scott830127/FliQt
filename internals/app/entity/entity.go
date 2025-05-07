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
	// LeaveQuota：先查員工與假別後批次建立
	var employees []Employee
	var leaveTypes []LeaveType
	if err := db.Find(&employees).Error; err != nil {
		return err
	}
	if err := db.Find(&leaveTypes).Error; err != nil {
		return err
	}
	type quotaEntry struct {
		LeaveTypeID uint
		Hours       int
	}
	// 建立對應的配額邏輯
	quotaMap := map[string]quotaEntry{
		"特休": {LeaveTypeID: 1, Hours: 80},
		"病假": {LeaveTypeID: 2, Hours: 40},
	}
	var quotas []LeaveQuota
	now := time.Now()
	for _, emp := range employees {
		for _, lt := range leaveTypes {
			if entry, ok := quotaMap[lt.Name]; ok {
				quotas = append(quotas, LeaveQuota{
					ID:          util.GetUniqueID(),
					EmployeeID:  emp.ID,
					LeaveTypeID: entry.LeaveTypeID,
					TotalHours:  entry.Hours,
					UsedHours:   0,
					UpdatedAt:   now,
				})
			}
		}
	}
	// 避免重複建立
	db.Model(&LeaveQuota{}).Count(&count)
	if count == 0 {
		if err := db.Create(&quotas).Error; err != nil {
			return err
		}
	}
	return nil
}
