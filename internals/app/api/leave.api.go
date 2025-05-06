package api

import (
	"FliQt/internals/app/dto"
	"FliQt/internals/app/service"
	"FliQt/pkg/redisx"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ ILeaveAPI = (*LeaveAPI)(nil)

type ILeaveAPI interface {
	Create(c *gin.Context)
	QueryTypes(c *gin.Context)
}

type LeaveAPI struct {
	srv   *service.Srv
	redis *redisx.Bundle
}

func NewLeaveAPI(srv *service.Srv, redis *redisx.Bundle) ILeaveAPI {
	return &LeaveAPI{srv: srv, redis: redis}
}

func (a *LeaveAPI) Create(c *gin.Context) {
	var params dto.LeaveRecordCreateCommand
	if err := c.ShouldBindJSON(params); err != nil {
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
	if err = a.srv.LeaveService.Create(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (a *LeaveAPI) QueryTypes(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := a.srv.LeaveService.QueryTypes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
