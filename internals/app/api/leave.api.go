package api

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/service"
	"FliQt/pkg/redisx"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var _ ILeaveAPI = (*LeaveAPI)(nil)

type ILeaveAPI interface {
	CreateRecord(c *gin.Context)      // member送假單
	QueryQuota(c *gin.Context)        // member查詢請假狀況
	AdminQueryRecord(c *gin.Context)  // admin查詢待審核假單
	AdminUpdateRecord(c *gin.Context) // admin審核
	QueryTypes(c *gin.Context)
}

type LeaveAPI struct {
	srv   *service.Srv
	redis *redisx.Bundle
}

func NewLeaveAPI(srv *service.Srv, redis *redisx.Bundle) ILeaveAPI {
	return &LeaveAPI{srv: srv, redis: redis}
}

func (a *LeaveAPI) CreateRecord(c *gin.Context) {
	var params dto.LeaveRecordCreateCommand
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	key := fmt.Sprintf("lock_leave_%d", params.EmployeeID)
	unlock, err := a.redis.Locker.AntiReSubmit(ctx, key)
	defer unlock()
	if err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		return
	}
	if err = a.srv.LeaveService.CreateRecord(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (a *LeaveAPI) QueryQuota(c *gin.Context) {
	employeeID, err := strconv.ParseUint(c.Query("employeeID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employeeID"})
		return
	}
	ctx := c.Request.Context()
	results, err := a.srv.LeaveService.QueryQuota(ctx, employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (a *LeaveAPI) AdminQueryRecord(c *gin.Context) {
	var query dto.LeaveRecordQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	results, err := a.srv.LeaveService.AdminQueryRecord(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (a *LeaveAPI) AdminUpdateRecord(c *gin.Context) {
	var cmd dto.LeaveRecordUpdateCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	if err := a.srv.LeaveService.AdminUpdateRecordAndQuota(ctx, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (a *LeaveAPI) QueryTypes(c *gin.Context) {
	ctx := c.Request.Context()
	results, err := a.srv.LeaveService.QueryTypes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
