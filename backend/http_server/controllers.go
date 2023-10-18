package http_server

import (
	"botdetector/app"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type controllers struct{}

func (controllers) Write(c *gin.Context) {
	data := app.RequestData{}

	if err := c.ShouldBindWith(&data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (controllers) Read(c *gin.Context) {
	data := app.RequestData{
		PlayerId:  123,
		Ip:        net.IPv4(123, 144, 111, 123),
		UserAgent: "teste",
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, data)
}

var Controllers = controllers{}
