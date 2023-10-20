package server

import (
	"botdetector/config"
	"botdetector/data"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	server := gin.Default()

	data.Database.Connect(config.Env.DbUri)

	go data.AsyncWriter.Start()

	server.GET("/logs", Controllers.ReadLogs)
	server.POST("/logs", Controllers.SaveLog)

	err := server.Run("0.0.0.0:" + config.Env.Port)

	if err != nil {
		log.Fatal("Can't start the http server: " + err.Error())
	}
}
