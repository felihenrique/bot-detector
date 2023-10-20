package server

import (
	"botdetector/app"
	"botdetector/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type controllers struct{}

func (controllers) SaveLog(c *gin.Context) {
	var item domain.RequestLog

	if err := c.ShouldBindWith(&item, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.Services.WriteLog(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, item)
}

type readLogsParams struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
	PlayerId  int    `form:"player_id"`
}

func (controllers) ReadLogs(c *gin.Context) {
	var params readLogsParams

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := app.Services.ReadLogs(app.ReadLogsParams(params))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

var Controllers = controllers{}
