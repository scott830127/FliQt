package router

import (
	"FliQt/internals/app/api"
	"FliQt/internals/app/config"
	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config, leaveAPI api.ILeaveAPI) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	fliqt := r.Group("/fliqt")
	{
		fliqt.POST("/leave", leaveAPI.Create)
	}

	return r
}
