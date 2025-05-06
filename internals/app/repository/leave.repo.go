package repository

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/entity"
	"FliQt/pkg/redisx"
	"FliQt/pkg/util"
	"context"
	"gorm.io/gorm"
	"time"
)

var _ ILeaveRepository = (*LeaveRepository)(nil)

type ILeaveRepository interface {
	Create(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error
	QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error)
}

type LeaveRepository struct {
	redis *redisx.Bundle
	db    *gorm.DB
}

func NewLeaveRepository(redis *redisx.Bundle, db *gorm.DB) ILeaveRepository {
	return &LeaveRepository{redis: redis, db: db}
}

func (r *LeaveRepository) Create(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error {
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
