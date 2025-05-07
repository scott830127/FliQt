package repository

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/entity"
	"FliQt/internals/app/enum"
	"FliQt/pkg/redisx"
	"FliQt/pkg/util"
	"context"
	"gorm.io/gorm"
	"time"
)

var _ ILeaveRepository = (*LeaveRepository)(nil)

type ILeaveRepository interface {
	CreateRecord(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error
	QueryQuota(ctx context.Context, query dto.LeaveQuotaQuery) ([]*dto.LeaveQuotaResult, error)
	QueryRecord(ctx context.Context, query dto.LeaveRecordQuery) ([]*dto.LeaveRecordResult, error)
	AdminUpdateRecordAndQuota(ctx context.Context, cmd dto.LeaveRecordUpdateCommand) error
	QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error)
}

type LeaveRepository struct {
	redis *redisx.Bundle
	db    *gorm.DB
}

func NewLeaveRepository(redis *redisx.Bundle, db *gorm.DB) ILeaveRepository {
	return &LeaveRepository{redis: redis, db: db}
}

func (r *LeaveRepository) CreateRecord(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error {
	record := entity.LeaveRecord{
		ID:          cmd.ID,
		EmployeeID:  cmd.EmployeeID,
		LeaveTypeID: cmd.LeaveTypeID,
		StartTime:   cmd.StartTime,
		EndTime:     cmd.EndTime,
		Hours:       cmd.Hours,
		Reason:      cmd.Reason,
		Status:      cmd.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return r.db.WithContext(ctx).Create(&record).Error
}

func (r *LeaveRepository) QueryQuota(ctx context.Context, query dto.LeaveQuotaQuery) ([]*dto.LeaveQuotaResult, error) {
	db := r.db.WithContext(ctx).
		Preload("Employee").
		Preload("LeaveType")

	if query.EmployeeID > 0 {
		db = db.Where("employee_id = ?", query.EmployeeID)
	}
	if query.LeaveTypeID > 0 {
		db = db.Where("leave_type_id = ?", query.LeaveTypeID)
	}
	var quotas []entity.LeaveQuota
	if err := db.Find(&quotas).Error; err != nil {
		return nil, err
	}

	var results []*dto.LeaveQuotaResult
	for _, q := range quotas {
		results = append(results, &dto.LeaveQuotaResult{
			EmployeeID:    q.EmployeeID,
			EmployeeName:  q.Employee.Name,
			Position:      q.Employee.Position,
			Email:         q.Employee.Email,
			Salary:        q.Employee.Salary,
			LeaveTypeID:   q.LeaveTypeID,
			LeaveTypeName: q.LeaveType.Name,
			Description:   q.LeaveType.Description,
			TotalHours:    q.TotalHours,
			UsedHours:     q.UsedHours,
			UpdatedAt:     q.UpdatedAt,
		})
	}
	return results, nil
}

func (r *LeaveRepository) QueryRecord(ctx context.Context, query dto.LeaveRecordQuery) ([]*dto.LeaveRecordResult, error) {
	db := r.db.WithContext(ctx).Model(&entity.LeaveRecord{}).
		Preload("Employee").
		Preload("LeaveType")

	if query.RecordID > 0 {
		db = db.Where("id = ?", query.RecordID)
	}
	if query.EmployeeID > 0 {
		db = db.Where("employee_id = ?", query.EmployeeID)
	}
	if query.LeaveTypeID > 0 {
		db = db.Where("leave_type_id = ?", query.LeaveTypeID)
	}
	if query.Hours > 0 {
		db = db.Where("hours < ?", query.Hours)
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}

	var records []entity.LeaveRecord
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}

	var results []*dto.LeaveRecordResult
	for _, r := range records {
		results = append(results, &dto.LeaveRecordResult{
			RecordID:      r.ID,
			EmployeeID:    r.EmployeeID,
			EmployeeName:  r.Employee.Name,
			Position:      r.Employee.Position,
			Email:         r.Employee.Email,
			Salary:        r.Employee.Salary,
			LeaveTypeID:   r.LeaveTypeID,
			LeaveTypeName: r.LeaveType.Name,
			Description:   r.LeaveType.Description,
			StartTime:     r.StartTime,
			EndTime:       r.EndTime,
			Hours:         r.Hours,
			Reason:        r.Reason,
			Status:        r.Status,
			CreatedAt:     r.CreatedAt,
			UpdatedAt:     r.UpdatedAt,
		})
	}
	return results, nil
}

func (r *LeaveRepository) AdminUpdateRecordAndQuota(ctx context.Context, cmd dto.LeaveRecordUpdateCommand) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var query dto.LeaveRecordQuery
		query.RecordID = cmd.RecordID
		records, err := r.QueryRecord(ctx, query)
		if err != nil {
			return err
		} else if len(records) == 0 {
			return nil
		}

		if err = tx.Model(&entity.LeaveRecord{}).
			Where("id = ?", cmd.RecordID).
			Update("status", cmd.Status).Error; err != nil {
			return err
		}

		if cmd.Status == enum.LeaveRecordStatusAccepted {
			if err = tx.Model(&entity.LeaveQuota{}).
				Where("employee_id = ? AND leave_type_id = ?", records[0].EmployeeID, records[0].LeaveTypeID).
				Update("used_hours", gorm.Expr("used_hours + ?", records[0].Hours)).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *LeaveRepository) AdminUpdateQuota(ctx context.Context, cmd dto.LeaveQuotaUpdateCommand) error {
	return r.db.WithContext(ctx).
		Model(&entity.LeaveQuota{}).
		Where("employee_id = ?", cmd.EmployeeID).
		Where("leave_type_id = ?", cmd.LeaveTypeID).
		Update("used_hours", cmd.UsedHours).Error
}

func (r *LeaveRepository) QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error) {
	const redisKey = "leave_type"
	var results []*dto.LeaveTypeResult
	if err := r.redis.Client.Get(ctx, redisKey).Scan(&results); err == nil && len(results) > 0 {
		return results, nil
	}
	var data []entity.LeaveType
	if err := r.db.WithContext(ctx).Find(&data).Error; err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []*dto.LeaveTypeResult{}, nil
	}
	for _, item := range data {
		resultItem := new(dto.LeaveTypeResult)
		if err := util.Copy(resultItem, item); err != nil {
			println("copy error")
		}
		results = append(results, resultItem)
	}
	_ = r.redis.Client.Set(ctx, redisKey, results, time.Hour) // 快取全部請假類型
	return results, nil
}
