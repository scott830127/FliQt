package service

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/repository"
)

var _ ILeaveService = (*LeaveService)(nil)

type ILeaveService interface {
	ApplyLeave(req dto.LeaveRequest) error
}

type LeaveService struct {
	repo *repository.Repo
}

func NewLeaveService(r *repository.Repo) ILeaveService {
	return &LeaveService{repo: r}
}

func (s *LeaveService) ApplyLeave(req dto.LeaveRequest) error {
	return s.repo.LeaveRepository.InsertLeave(req.EmployeeID, req.Reason, req.Days)
}
