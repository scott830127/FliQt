package service

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/enum"
	"FliQt/internals/app/repository"
	"FliQt/pkg/util"
	"context"
)

var _ ILeaveService = (*LeaveService)(nil)

type ILeaveService interface {
	Create(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error
	QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error)
}

type LeaveService struct {
	repo *repository.Repo
}

func NewLeaveService(r *repository.Repo) ILeaveService {
	return &LeaveService{repo: r}
}

func (s *LeaveService) Create(ctx context.Context, cmd dto.LeaveRecordCreateCommand) error {
	cmd.ID = util.GetUniqueID()
	cmd.Status = enum.LeaveRecordStatusPending
	return s.repo.LeaveRepository.Create(ctx, cmd)
}

func (s *LeaveService) QueryTypes(ctx context.Context) ([]*dto.LeaveTypeResult, error) {
	return s.repo.LeaveRepository.QueryTypes(ctx)
}
