package router

import (
	"FliQt/internals/app/api"
	"github.com/gin-gonic/gin"
)

func New(leaveAPI api.ILeaveAPI) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	fliqt := r.Group("/fliqt")
	{
		fliqt.POST("/member/leave/record", leaveAPI.CreateRecord)
		fliqt.GET("/member/leave/quota", leaveAPI.QueryQuota)
		fliqt.GET("/member/leave/type", leaveAPI.QueryTypes)

		fliqt.GET("/admin/leave/record", leaveAPI.AdminQueryRecord)
		fliqt.PATCH("/admin/leave/record", leaveAPI.AdminUpdateRecord)
	}
	return r
}
