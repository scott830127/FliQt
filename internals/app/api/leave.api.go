package api

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ ILeaveAPI = (*LeaveAPI)(nil)

type ILeaveAPI interface {
	Apply(c *gin.Context)
	Create(c *gin.Context)
}

type LeaveAPI struct {
	srv *service.Srv
}

func NewLeaveAPI(srv *service.Srv) ILeaveAPI {
	return &LeaveAPI{srv: srv}
}

func (a *LeaveAPI) Apply(c *gin.Context) {
	var req dto.LeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.srv.LeaveService.ApplyLeave(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "leave request submitted"})
}

func (a *LeaveAPI) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "leave request submitted"})
	return
}
