package server

import (
	"botdetector/app"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type controllers struct{}

func (controllers) SaveRequest(c *gin.Context) {
	var data app.RequestData

	if err := c.ShouldBindWith(&data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.Services.HydrateRequestData(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (controllers) ReadRequests(c *gin.Context) {
	var query struct {
		ChannelId string `form:"channel_id" binding:"required"`
		Start     string `form:"start" binding:"required"`
		End       string `form:"end" binding:"required"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := app.RequestData{
		PlayerId:  123,
		Ip:        net.IPv4(123, 144, 111, 123),
		UserAgent: "teste",
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, data)
}

var Controllers = controllers{}
