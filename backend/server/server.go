package server

import (
	"botdetector/app"
	"botdetector/config"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	server := gin.Default()

	server.GET("/requests", Controllers.Read)
	server.POST("/requests", Controllers.Write)

	app.Detector.LoadIps(config.Env.IpsFilePath)

	err := server.Run("0.0.0.0:" + config.Env.Port)

	if err != nil {
		log.Fatal("Can't start the http server: " + err.Error())
	}
}
