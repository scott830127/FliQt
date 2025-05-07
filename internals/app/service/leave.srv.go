package service

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/enum"
	"FliQt/internals/app/repository"
	"FliQt/pkg/util"
	"context"
	"fmt"
	"sync"
)

var _ ILeaveService = (*LeaveService)(nil)

type ILeaveService interface {
	CreateRecord(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error
	QueryQuota(ctx context.Context, employeeID uint64) ([]*dto.LeaveQuotaResult, error)
	AdminQueryRecord(ctx context.Context, query dto.LeaveRecordQuery) ([]*dto.LeaveRecordResult, error)
	AdminUpdateRecordAndQuota(ctx context.Context, cmd dto.LeaveRecordUpdateCommand) error
	QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error)
}

type LeaveService struct {
	repo *repository.Repo
}

func NewLeaveService(r *repository.Repo) ILeaveService {
	return &LeaveService{repo: r}
}

func (s *LeaveService) CreateRecord(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error {
	//取得剩餘時數與待審核時數加總若超過則不給請假
	var quotas []*dto.LeaveQuotaResult
	var records []*dto.LeaveRecordResult
	var quotaErr, recordErr error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var queryQuota dto.LeaveQuotaQuery
		queryQuota.EmployeeID = cmd.EmployeeID
		queryQuota.LeaveTypeID = cmd.LeaveTypeID
		quotas, quotaErr = s.repo.LeaveRepository.QueryQuota(ctx, queryQuota)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var queryRecord dto.LeaveRecordQuery
		queryRecord.EmployeeID = cmd.EmployeeID
		queryRecord.LeaveTypeID = cmd.LeaveTypeID
		queryRecord.Status = enum.LeaveRecordStatusPending
		records, recordErr = s.repo.LeaveRepository.QueryRecord(ctx, queryRecord)
	}()
	wg.Wait()
	if quotaErr != nil {
		return quotaErr
	}
	if recordErr != nil {
		return recordErr
	}

	var applyHours int
	for _, record := range records {
		applyHours += record.Hours
	}
	if applyHours+quotas[0].UsedHours+cmd.Hours > quotas[0].TotalHours {
		return fmt.Errorf("is over total hours")
	}

	cmd.ID = util.GetUniqueID()
	cmd.Status = enum.LeaveRecordStatusPending
	return s.repo.LeaveRepository.CreateRecord(ctx, cmd)
}

func (s *LeaveService) QueryQuota(ctx context.Context, employeeID uint64) ([]*dto.LeaveQuotaResult, error) {
	var query dto.LeaveQuotaQuery
	query.EmployeeID = employeeID
	return s.repo.LeaveRepository.QueryQuota(ctx, query)
}

func (s *LeaveService) AdminQueryRecord(ctx context.Context, query dto.LeaveRecordQuery) ([]*dto.LeaveRecordResult, error) {
	return s.repo.LeaveRepository.QueryRecord(ctx, query)
}

func (s *LeaveService) AdminUpdateRecordAndQuota(ctx context.Context, cmd dto.LeaveRecordUpdateCommand) error {
	if cmd.Status <= 1 {
		return nil
	}
	return s.repo.LeaveRepository.AdminUpdateRecordAndQuota(ctx, cmd)
}

func (s *LeaveService) QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error) {
	return s.repo.LeaveRepository.QueryTypes(ctx)
}
